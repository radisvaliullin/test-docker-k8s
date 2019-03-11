package testsrv

import (
	"context"
	"log"
	"time"

	"github.com/radisvaliullin/test-docker-k8s/proto/api/v1/tsrv"
)

// Req -
func (s *Server) Req(ctx context.Context, req *tsrv.Request) (*tsrv.Response, error) {
	log.Printf("req start: id - %v, msg - %v", req.GetId(), req.GetMsg())
	resp := &tsrv.Response{Id: req.GetId(), Msg: req.GetMsg()}
	log.Printf("req end: id - %v, msg - %v", req.GetId(), req.GetMsg())
	return resp, nil
}

// ReqDelay -
func (s *Server) ReqDelay(ctx context.Context, req *tsrv.Request) (*tsrv.Response, error) {
	log.Printf("reqDelay start: id - %v, msg - %v", req.GetId(), req.GetMsg())
	time.Sleep(time.Millisecond * 10000)
	resp := &tsrv.Response{Id: req.GetId(), Msg: req.GetMsg()}
	log.Printf("reqDelay end: id - %v, msg - %v", req.GetId(), req.GetMsg())
	return resp, nil
}

// Req2 -
func (s *Server) Req2(req *tsrv.Request, stream tsrv.TestService_Req2Server) error {
	log.Printf("req2 start: id - %v, msg - %v", req.GetId(), req.GetMsg())
	//resp := &tsrv.Response{Id: req.GetId(), Msg: req.GetMsg()}
	log.Printf("req2 end: id - %v, msg - %v", req.GetId(), req.GetMsg())
	return nil
}

// Req3 -
func (s *Server) Req3(stream tsrv.TestService_Req3Server) error {
	// log.Printf("req3 start: id - %v, msg - %v", req.GetId(), req.GetMsg())
	// resp := &tsrv.Response{Id: req.GetId(), Msg: req.GetMsg()}
	// log.Printf("req3 end: id - %v, msg - %v", req.GetId(), req.GetMsg())
	return nil
}

// Req4 -
func (s *Server) Req4(stream tsrv.TestService_Req4Server) error {
	// log.Printf("req4 start: id - %v, msg - %v", req.GetId(), req.GetMsg())
	// resp := &tsrv.Response{Id: req.GetId(), Msg: req.GetMsg()}
	// log.Printf("req4 end: id - %v, msg - %v", req.GetId(), req.GetMsg())
	return nil
}
