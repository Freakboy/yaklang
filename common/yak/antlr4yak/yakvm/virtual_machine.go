package yakvm

import (
	"context"
	"fmt"
	"github.com/yaklang/yaklang/common/log"
	"github.com/yaklang/yaklang/common/utils"
	"github.com/yaklang/yaklang/common/yak/antlr4yak/yakvm/vmstack"
	"sync"
)

type ExecFlag int

const (
	None   ExecFlag = 1 << iota // 默认创建新栈帧执行代码，执行后出栈
	Trace                       // 执行后不出站
	Sub                         // 使用栈顶的数据继续执行
	Inline                      // 使用上一次执行的Trace继续执行
)

func GetFlag(flags ...ExecFlag) ExecFlag {
	flag := None
	for _, f := range flags {
		flag |= f
	}
	return flag
}

type YakitFeedbacker interface {
}
type BreakPointFactoryFun func(v *VirtualMachine) bool
type VirtualMachine struct {
	//globalVar 是当前引擎的全局变量，属于引擎
	globalVar map[string]interface{}
	VMStack   *vmstack.Stack
	lastPanic *VMPanic
	rootScope *Scope

	// debug
	debug           bool // 内部debug
	debugMode       bool // 外部debugger
	debugger        *Debugger
	BreakPoint      []BreakPointFactoryFun
	yakitFeedbacker YakitFeedbacker
	config          *VirtualMachineConfig
	// map[sha1(caller, callee)]func(any)any
	hijackMapMemberCallHandlers sync.Map
	globalVarFallback           func(string) interface{}
}

func (n *VirtualMachine) RegisterMapMemberCallHandler(caller, callee string, h func(interface{}) interface{}) {
	n.hijackMapMemberCallHandlers.Store(utils.CalcSha1(caller, callee), h)
}

func (n *VirtualMachine) RegisterGlobalVariableFallback(h func(string) interface{}) {
	n.globalVarFallback = h
}

func (n *VirtualMachine) SetYaiktFeedbacker(i YakitFeedbacker) {
	n.yakitFeedbacker = i
}
func (v *VirtualMachine) AddBreakPoint(fun BreakPointFactoryFun) {
	v.BreakPoint = append(v.BreakPoint, fun)
}
func (n *VirtualMachine) GetExternalVariableNames() []string {
	vs := []string{}
	for k, _ := range n.globalVar {
		vs = append(vs, k)
	}
	return vs
}
func (v *VirtualMachine) SetDebug(debug bool) {
	v.debug = debug
}
func (v *VirtualMachine) SetDebugMode(debug bool, sourceCode string, codes []*Code, debugInit, debugCallback func(*Debugger)) {
	v.debugMode = debug
	if !debug {
		return
	}
	if v.debugger == nil {
		v.debugger = NewDebugger(v, sourceCode, codes, debugInit, debugCallback)
	}
}

func (v *VirtualMachine) SetSymboltable(table *SymbolTable) {
	v.rootScope = NewScope(table)
}

func (v *VirtualMachine) panic(i *VMPanic) {
	v.lastPanic = i
}
func (v *VirtualMachine) recover() *VMPanic {
	lastPanic := v.lastPanic
	v.lastPanic = nil
	return lastPanic
}
func NewWithSymbolTable(table *SymbolTable) *VirtualMachine {
	v := &VirtualMachine{
		//rootSymbol: table,
		rootScope: NewScope(table),
		VMStack:   vmstack.New(),
		globalVar: make(map[string]interface{}),
		config:    NewVMConfig(),
	}
	return v
}
func New() *VirtualMachine {
	return NewWithSymbolTable(NewSymbolTable())
}

// deepCopyLib 拷贝yaklang依赖，防止多个engine并行运行时对lib进行hook导致concurrent write map error
func deepCopyLib(libs map[string]interface{}) map[string]interface{} {
	newLib := map[string]interface{}{}
	for k, v := range libs {
		if v1, ok := v.(map[string]interface{}); ok {
			newLib[k] = deepCopyLib(v1)
		} else {
			newLib[k] = v
		}
	}
	return newLib
}

// ImportLibs 导入库到引擎的全局变量中
func (n *VirtualMachine) ImportLibs(libs map[string]interface{}) {
	for k, v := range deepCopyLib(libs) {
		n.globalVar[k] = v
	}
}

// SetVar 导入变量到引擎的全局变量中
func (n *VirtualMachine) SetVar(k string, v interface{}) {
	n.globalVar[k] = v
}
func (n *VirtualMachine) GetVar(name string) (interface{}, bool) {
	ivm := n.VMStack.Peek()
	if ivm == nil {
		val, ok := n.rootScope.GetValueByName(name)
		if ok {
			return val.Value, true
		}
	} else {
		// ivm 存在的时候，从 frame 中找变量
		val, ok := ivm.(*Frame).CurrentScope().GetValueByName(name)
		if ok {
			return val.Value, true
		}
	}
	// 如果不存在，根也找不到，那么就直接全局找
	var_, ok := n.globalVar[name]
	if ok {
		return var_, true
	}

	if n.globalVarFallback != nil {
		hijackedGlobal := n.globalVarFallback(name)
		if hijackedGlobal != nil {
			return hijackedGlobal, true
		}
	}

	return undefined, false
}
func (n *VirtualMachine) GetGlobalVar() map[string]interface{} {
	return n.globalVar
}

func (v *VirtualMachine) ExecYakFunction(ctx context.Context, f *Function, args map[int]*Value, flags ...ExecFlag) (interface{}, error) {
	var value interface{}
	var finalFlags = []ExecFlag{Sub}
	if len(flags) > 0 {
		finalFlags = flags
	}
	err := v.Exec(ctx, func(frame *Frame) {
		name := f.GetActualName()
		frame.SetVerbose(fmt.Sprintf("function: %s", name))
		frame.SetFunction(f)
		if f.sourceCode != "" {
			frame.SetOriginCode(f.sourceCode)
		}
		//闭包继承父作用域
		//if v.config.GetClosureSupport() {
		frame.scope = f.scope
		frame.CreateAndSwitchSubScope(f.symbolTable)
		//} else {
		//	frame.scope = NewScope(f.symbolTable)
		//}
		for id, arg := range args {
			frame.CurrentScope().NewValueByID(id, arg)
		}
		frame.Exec(f.codes)
		if frame.lastStackValue != nil {
			value = frame.lastStackValue.Value
		}
		//if v.config.GetClosureSupport() {
		frame.ExitScope()
		//}
	}, finalFlags...)
	if err != nil {
		return nil, err
	}
	return value, nil
}
func (v *VirtualMachine) ExecAsyncYakFunction(ctx context.Context, f *Function, args map[int]*Value) error {
	return v.Exec(ctx, func(frame *Frame) {
		name := f.GetActualName()
		frame.SetVerbose("function: " + name)
		frame.SetFunction(f)
		frame.SetScope(f.scope)
		frame.CreateAndSwitchSubScope(f.symbolTable)
		for id, arg := range args {
			frame.CurrentScope().NewValueByID(id, arg)
		}
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("yakvm async function panic: %v", err)
					utils.PrintCurrentGoroutineRuntimeStack()
				}
			}()
			frame.Exec(f.codes)
			frame.ExitScope()
		}()

	}, Sub)
}

func (v *VirtualMachine) ExecYakCode(ctx context.Context, sourceCode string, codes []*Code, flags ...ExecFlag) error {
	return v.Exec(ctx, func(frame *Frame) {
		frame.SetVerbose("global code")
		frame.SetOriginCode(sourceCode)
		frame.Exec(codes)
	}, flags...)
}
func (v *VirtualMachine) InlineExecYakCode(ctx context.Context, codes []*Code, flags ...ExecFlag) error {
	return v.Exec(ctx, func(frame *Frame) {
		frame.Exec(codes)
	}, Trace|Sub)
}

var vmstackLock = new(sync.Mutex)

func (v *VirtualMachine) Exec(ctx context.Context, f func(frame *Frame), flags ...ExecFlag) error {
	flag := GetFlag(flags...)

	var frame *Frame
	if flag&Sub == Sub {

		vmstackLock.Lock()
		topFrame := v.VMStack.Peek()
		vmstackLock.Unlock()

		if topFrame == nil {
			panic("BUG: VMStack is empty(Sub)")
		}
		frame = NewSubFrame(topFrame.(*Frame))
	} else if flag&Inline == Inline {
		vmstackLock.Lock()
		topFrame := v.VMStack.Peek()
		vmstackLock.Unlock()

		if topFrame == nil {
			topFrame = NewFrame(v)
			vmstackLock.Lock()
			v.VMStack.Push(topFrame)
			vmstackLock.Unlock()
			log.Debugf("VMStack is empty(Inline), we create new frame")
		}

		frame = topFrame.(*Frame)
		codes := frame.codes
		p := frame.codePointer
		defer func() {
			frame.codes = codes
			frame.codePointer = p
		}()
	} else {
		frame = NewFrame(v)
		for k, val := range v.globalVar {
			frame.GlobalVariables[k] = val
		}
	}

	vmstackLock.Lock()
	v.VMStack.Push(frame)
	vmstackLock.Unlock()

	frame.debug = v.debug
	// 初始化debugger
	if v.debugMode && v.debugger != nil && v.debugger.initFunc != nil {
		v.debugger.InitCallBack()
	}
	frame.ctx = ctx
	//给eval函数注入context
	if v1, ok := frame.GlobalVariables["eval"]; ok {
		if v2, ok := v1.(func(ctx context.Context, code string)); ok {
			frame.GlobalVariables["eval"] = func(code string) {
				v2(ctx, code)
			}
		}
	}

	f(frame)

	//未设置Trace时执行后出站
	if flag&Trace != Trace {
		vmstackLock.Lock()
		v.VMStack.Pop()
		vmstackLock.Unlock()
	}
	if lastPanic := v.recover(); lastPanic != nil {
		lastPanic.contextInfos.Peek().(*PanicInfo).SetPostionVerbose(frame.GetVerbose())
		if exitValue, ok := lastPanic.data.(*VMPanicSignal); ok {
			panic(exitValue)
		} else {
			panic(lastPanic)
		}
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return nil
}

func (v *VirtualMachine) CurrentFM() *Frame {
	return v.VMStack.Peek().(*Frame)
}

func (v *VirtualMachine) GetConfig() *VirtualMachineConfig {
	return v.config
}

func (v *VirtualMachine) SetConfig(config *VirtualMachineConfig) {
	v.config = config
}
