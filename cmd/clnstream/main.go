package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/radisvaliullin/test-docker-k8s/proto/api/v1/tsrv"
	"google.golang.org/grpc"
)

var addr = "0.0.0.0:7373"

func main() {

	log.Printf("stream client, start %v", time.Now())

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("test srv cln dial err, %v", err)
	}
	defer conn.Close()

	// client
	cln := tsrv.NewTestServiceClient(conn)

	// make periodic requests

	// simple req/resp request
	resp, err := cln.Req(
		context.Background(),
		&tsrv.Request{Id: tnow(), Msg: "test req/resp request"},
	)
	if err != nil {
		log.Fatalf("stream client: simple req/resp request err - %v", err)
	}
	log.Printf("stream cln: simple request - %+v", resp)

	wg := sync.WaitGroup{}

	// // server side stream request
	// log.Printf("req srvSideStream1, %v", tnow())
	// wg.Add(1)
	// go srvSideStrm("srvSideStream1", cln, &wg)
	// log.Printf("req srvSideStream2, %v", tnow())
	// wg.Add(1)
	// srvSideStrm("ssrvSideStream2", cln, &wg)

	// wg.Wait()

	// client side stream request
	log.Printf("req clnSideStream1, %v", tnow())
	wg.Add(1)
	go clnSideStrm("clnSideStream1", cln, &wg)

	wg.Wait()
}

func clnSideStrm(pref string, cln tsrv.TestServiceClient, wg *sync.WaitGroup) {
	defer wg.Done()

	// streaming
	tn := tnow()
	strm, err := cln.ReqClnStream(context.Background())
	if err != nil {
		log.Fatalf("%v: err - %v, start - %v, now - %v", pref, err, tn, tnow())
	}
	log.Printf("%v requested, start - %v, now - %v", pref, tn, tnow())

	// sending
	for i := 0; i < 4; i++ {
		reqClnStrm := &tsrv.Request{
			Id:  tn,
			Msg: fmt.Sprintf("%v start - %v now - %v", pref, tn, tnow()),
		}
		if err := strm.Send(reqClnStrm); err != nil {
			if err == io.EOF {
				log.Printf(
					"%v EOF: req - %+v, start - %v, now - %v",
					pref, reqClnStrm, tn, tnow())
				break
			}
			log.Fatalf(
				"%v error: req - %+v, err - %v, start - %v, now - %v",
				pref, reqClnStrm, err, tn, tnow())
		}
		log.Printf("%v sent: req - %+v, start - %v, now - %v", pref, reqClnStrm, tn, tnow())
	}

	// reply
	reply, err := strm.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v closeAndRecv: err - %v, start - %v, now - %v", pref, err, tn, tnow())
	}
	log.Printf("%v sent: reply - %+v, start - %v, now - %v", pref, reply, tn, tnow())
}

func srvSideStrm(pref string, cln tsrv.TestServiceClient, wg *sync.WaitGroup) {
	defer wg.Done()

	// request
	tn := tnow()
	reqSrvStrm := &tsrv.Request{
		Id:  tn,
		Msg: fmt.Sprintf("%v %v", pref, tn),
	}
	strm, err := cln.ReqSrvStream(context.Background(), reqSrvStrm)
	if err != nil {
		log.Fatalf("%v: req - %+v, err - %v", pref, reqSrvStrm, err)
	}
	log.Printf("%v requested, now - %v", pref, tnow())

	// reciving
	wg.Add(1)
	go func() {
		defer wg.Done()

		time.Sleep(time.Millisecond * 1000)
		for {
			respSrvStrm, err := strm.Recv()
			if err == io.EOF {
				log.Printf("%v EQF: req - %+v, err - %v", pref, reqSrvStrm, err)
			}
			if err != nil {
				log.Fatalf("%v, recv err - %v", pref, err)
			}
			log.Printf("%v, recv resp - %+v, now - %v", pref, respSrvStrm, tnow())
		}
	}()
}

func tnow() int64 {
	return time.Now().UnixNano() / 1000000
}
