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
	hnow := tnow()
	log.Printf("reqSrvStream start: req - %+v, hnow - %v", req, hnow)

	for i := 0; i < 2; i++ {
		resp := &tsrv.Response{
			Id: req.GetId(),
			Msg: fmt.Sprintf(
				"reqMsg - %v, i - %v, hnow - %v, now - %v", req.GetMsg(), i, hnow, tnow()),
		}
		if err := stream.Send(resp); err != nil {
			log.Printf(
				"reqSrvStream err: req - %+v, i - %v, err - %v, hnow - %v, now - %v",
				req, i, err, hnow, tnow())
			return err
		}
		log.Printf(
			"reqSrvStream: req - %+v, resp - %+v, i - %v, hnow - %v, now - %v",
			req, resp, i, hnow, tnow())
		time.Sleep(time.Millisecond * 500)
	}

	log.Printf("reqSrvStream end: req - %+v, hnow - %v, now - %v", req, hnow, tnow())
	return nil
}

// ReqClnStream -
func (s *Server) ReqClnStream(stream tsrv.TestService_ReqClnStreamServer) error {
	hnow := tnow()
	log.Printf("reqClnStream start: hnow - %v", hnow)

	for {
		// recive
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("reqClnStream EOF - %+v, hnow - %v, now - %v", req, hnow, tnow())
			return stream.SendAndClose(&tsrv.Response{
				Id: req.GetId(),
				Msg: fmt.Sprintf("reqMsg - %v, hnow - %v, now - %v",
					req.GetMsg(), hnow, tnow()),
			})
		}
		if err != nil {
			log.Printf("reqClnStream err - %v, hnow - %v", err, hnow)
			return err
		}
		time.Sleep(time.Millisecond * 1000)
		log.Printf("reqClnStream accept: req - %+v, hnow - %v, now - %v", req, hnow, tnow())
	}
}

// ReqBiStream -
func (s *Server) ReqBiStream(stream tsrv.TestService_ReqBiStreamServer) error {
	hnow := tnow()
	log.Printf("reqBiStream, start - %v", hnow)

	for {
		// reciving
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("reqBiStream EOF, hnow - %v, now - %v", hnow, tnow())
			return nil
		}
		if err != nil {
			log.Printf("reqBiStream error, err - %v, hnow - %v, now - %v", err, hnow, tnow())
			return err
		}
		log.Printf("reqBiStream recived, req - %+v, hnow - %v, now - %v",
			req, hnow, tnow())

		// sending
		time.Sleep(time.Millisecond * 500)
		log.Printf("reqBiStream sending, hnow - %v, now - %v", hnow, tnow())
		for i := 0; i < 2; i++ {
			resp := &tsrv.Response{
				Id: req.GetId(),
				Msg: fmt.Sprintf("reqMsg - %+v i - %v hnow - %v, now - %v",
					req.GetMsg(), i, hnow, tnow()),
			}
			if err := stream.Send(resp); err != nil {
				log.Printf(
					"reqBiStream send err: resp - %+v, i - %v, err - %v, hnow - %v, now - %v",
					resp, i, err, hnow, tnow())
				return err
			}
			log.Printf("reqBiStream sent, resp - %+v, hnow - %v, now - %v",
				resp, hnow, tnow())
		}
	}
}

func tnow() int64 {
	return time.Now().UnixNano() / 1000000
}
