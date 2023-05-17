package yakgrpc

import (
	"context"
	"github.com/yaklang/yaklang/common/consts"
	"github.com/yaklang/yaklang/common/utils"
	"github.com/yaklang/yaklang/common/yakgrpc/ypb"
	"strings"
	"sync"
)

var (
	isInAttached         = utils.NewBool(false)
	attachOutputCallback = new(sync.Map)
)

func (s *Server) AttachCombinedOutput(req *ypb.AttachCombinedOutputRequest, server ypb.AttachApi_AttachCombinedOutputServer) error {
	lines := utils.GetCachedLog()
	server.Send(&ypb.ExecResult{
		Raw: []byte("===========History log=========\n" + strings.Join(lines, "\n") + "\n===========Real time log=========\n"),
	})
	return utils.HandleStdout(server.Context(), func(s string) {
		server.Send(&ypb.ExecResult{
			Raw: []byte(s),
		})
	})
}

func (s *Server) SaveTextToTemporalFile(ctx context.Context, req *ypb.SaveTextToTemporalFileRequest) (*ypb.SaveTextToTemporalFileResponse, error) {
	fileName, err := consts.TempFile("tmp*.txt")
	if err != nil {
		return nil, err
	}
	fileName.Write(req.GetText())
	fileName.Close()
	return &ypb.SaveTextToTemporalFileResponse{FileName: fileName.Name()}, nil
}
