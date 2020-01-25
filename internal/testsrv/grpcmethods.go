package testsrv

import (
	"context"
	"log"

	"github.com/radisvaliullin/test-docker-k8s/proto/api/v1/testsrv"
)

// Req -
func (s *Server) Req(ctx context.Context, req *testsrv.Request) (*testsrv.Response, error) {
	resp := &testsrv.Response{Id: req.GetId(), Msg: req.GetMsg()}
	log.Printf("req: id - %v, msg - %v", req.GetId(), req.GetMsg())
	return resp, nil
}
