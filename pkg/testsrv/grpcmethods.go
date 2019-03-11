package testsrv

import (
	"context"
	"fmt"
	"io"
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
	for i := 0; i < 10; i++ {
		resp := &tsrv.Response{Id: req.GetId(), Msg: fmt.Sprintf("%v : i - %v", req.GetMsg(), i)}
		if err := stream.Send(resp); err != nil {
			log.Printf("req2 err: id - %v, msg - %v, i - %v, err - %v", req.GetId(), req.GetMsg(), i, err)
			return err
		}
	}
	log.Printf("req2 end: id - %v, msg - %v", req.GetId(), req.GetMsg())
	return nil
}

// Req3 -
func (s *Server) Req3(stream tsrv.TestService_Req3Server) error {
	log.Printf("req3 start")

	for {
		// recive
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("req3 EOF - %+v", req)
			return stream.SendAndClose(&tsrv.Response{
				Id:  req.GetId(),
				Msg: req.GetMsg(),
			})
		}
		if err != nil {
			log.Printf("req3 err - %v", err)
			return err
		}

		log.Printf("req3 accept: req - %+v", req)
	}

	// log.Printf("req3 end")
	// return nil
}

// Req4 -
func (s *Server) Req4(stream tsrv.TestService_Req4Server) error {
	log.Printf("req4 start")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("req4 EOF - %+v", req)
			return nil
		}
		if err != nil {
			log.Printf("req4 err - %+v", err)
			return err
		}

		for i := 0; i < 10; i++ {
			resp := &tsrv.Response{Id: req.GetId(), Msg: fmt.Sprintf("%v : i - %v", req.GetMsg(), i)}
			if err := stream.Send(resp); err != nil {
				log.Printf("req4 err: id - %v, msg - %v, i - %v, err - %v", req.GetId(), req.GetMsg(), i, err)
				return err
			}
		}
	}

	// log.Printf("req4 end: id - %v, msg - %v", req.GetId(), req.GetMsg())
	// return nil
}
