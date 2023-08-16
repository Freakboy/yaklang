package suricata

import (
	"math"
)

type Generator interface {
	Gen() []byte
}

type Surigen struct {
	rules []*ContentRule
	gen   map[Modifier]Generator
}

func NewSurigen(contentRules []*ContentRule) (*Surigen, error) {
	g := &Surigen{
		rules: contentRules,
		gen:   make(map[Modifier]Generator),
	}
	err := g.parse()
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Surigen) Gen() ([]byte, error) {
	mp := make(map[Modifier][]byte)
	for k, gener := range g.gen {
		mp[k] = gener.Gen()
	}
	return HTTPCombination(mp), nil
}

func (g *Surigen) parse() error {
	// mapping by Modifier
	var mp = make(map[Modifier][]*ContentRule)
	for _, rule := range g.rules {
		mp[rule.Modifier] = append(mp[rule.Modifier], rule)
	}

	// parse rules
	for mdf, rule := range mp {
		// special part use special generator
		// designed but not in using tempetarily
		switch mdf {
		case HTTPStatCode:
			g.gen[mdf] = parse2ContentGen(rule, WithNoise(noiseDigit), WithTryLen(3))
		default:
			g.gen[mdf] = parse2ContentGen(rule)
		}
	}

	// set Len
	for _, payload := range g.gen {
		payload, ok := payload.(*ContentGen)
		if !ok {
			continue
		}
		for _, m := range payload.Modifiers {
			switch m := m.(type) {
			case *ContentModifier:
				if m.Offset+len(m.Content) > payload.Len || m.Relative {
					payload.Len += m.Offset + len(m.Content)
				}
			case *RegexpModifier:
				payload.Len += len(m.Generator.Generate())
			}
		}
		payload.Len = 1 << math.Ilogb(float64(payload.Len+1))
	}

	return nil
}
