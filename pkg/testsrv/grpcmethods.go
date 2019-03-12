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

// ReqSrvStream -
func (s *Server) ReqSrvStream(req *tsrv.Request, stream tsrv.TestService_ReqSrvStreamServer) error {
	now := time.Now()
	log.Printf("reqSrvStream start: req - %+v, %v", req, now)

	for i := 0; i < 4; i++ {
		resp := &tsrv.Response{
			Id:  req.GetId(),
			Msg: fmt.Sprintf("%v : i - %v, now - %v", req.GetMsg(), i, now),
		}
		if err := stream.Send(resp); err != nil {
			log.Printf("reqSrvStream err: req - %+v, i - %v, err - %v, %v", req, i, err, now)
			return err
		}
		time.Sleep(time.Millisecond * 500)
	}

	log.Printf("reqSrvStream end: req - %+v, %v", req, now)
	return nil
}

// ReqClnStream -
func (s *Server) ReqClnStream(stream tsrv.TestService_ReqClnStreamServer) error {
	now := time.Now()
	log.Printf("reqClnStream start: %v", now)

	for {
		// recive
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("reqClnStream EOF - %+v, %v", req, now)
			return stream.SendAndClose(&tsrv.Response{
				Id:  req.GetId(),
				Msg: req.GetMsg(),
			})
		}
		if err != nil {
			log.Printf("reqClnStream err - %v, %v", err, now)
			return err
		}

		log.Printf("reqClnStream accept: req - %+v, %v", req, now)
	}
}

// ReqBiStream -
func (s *Server) ReqBiStream(stream tsrv.TestService_ReqBiStreamServer) error {
	now := time.Now()
	log.Printf("reqBiStream start, %v", now)

	for {
		// accept
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("reqBiStream EOF - %+v, %v", req, now)
			return nil
		}
		if err != nil {
			log.Printf("reqBiStream err - %+v, %v", err, now)
			return err
		}

		// respond
		for i := 0; i < 4; i++ {
			resp := &tsrv.Response{
				Id:  req.GetId(),
				Msg: fmt.Sprintf("%v : i - %v, %v", req.GetMsg(), i, now),
			}
			if err := stream.Send(resp); err != nil {
				log.Printf("reqBiStream err: req - %+v, i - %v, err - %v, %v", req, i, err, now)
				return err
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
}
