package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/radisvaliullin/test-docker-k8s/proto/api/v1/tsrv"

	"google.golang.org/grpc"
)

var addr = "0.0.0.0:7373"

func main() {

	log.Printf("test server client start, addr - %v", addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("test srv cln dial err, %v", err)
	}
	defer conn.Close()

	// client
	cln := tsrv.NewTestServiceClient(conn)

	// make periodic requests

	// req
	go func() {

		for {
			tn := time.Now()
			r := &tsrv.Request{Id: tn.Unix(), Msg: fmt.Sprintf("%v", tn)}
			log.Printf("make req: req - %+v", r)

			resp, err := cln.Req(context.Background(), r)
			log.Printf("make req response: req - %+v, resp - %+v, err - %v", r, resp, err)

			time.Sleep(time.Millisecond * 1000)
		}
	}()

	// recDelay
	go func() {
		for {
			tn := time.Now()
			r := &tsrv.Request{Id: tn.Unix(), Msg: fmt.Sprintf("Delay %v", tn)}
			log.Printf("make recDelay: req - %+v", r)

			resp, err := cln.ReqDelay(context.Background(), r)
			log.Printf("make recDelay response: req - %+v, resp - %+v, err - %v", r, resp, err)

			time.Sleep(time.Millisecond * 1000)
		}
	}()

	// recDelay2
	for {
		tn := time.Now()
		r := &tsrv.Request{Id: tn.Unix(), Msg: fmt.Sprintf("Delay2 %v", tn)}
		log.Printf("make recDelay: req - %+v", r)

		resp, err := cln.ReqDelay(context.Background(), r)
		log.Printf("make recDelay response: req - %+v, resp - %+v, err - %v", r, resp, err)

		time.Sleep(time.Millisecond * 1000)
	}
}
