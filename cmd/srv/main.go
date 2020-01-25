package main

import (
	"log"

	"github.com/radisvaliullin/test-docker-k8s/internal/testsrv"
)

var gAddr = "0.0.0.0:7373"
var hAddr = "0.0.0.0:7374"

func main() {

	log.Printf("test server start, gAddr - %v, hAddr - %v", gAddr, hAddr)

	// init, start server
	conf := testsrv.Config{GAddr: gAddr, HAddr: hAddr}
	gSrv := testsrv.New(conf)
	err := gSrv.Start()
	if err != nil {
		log.Fatalf("test server start err, %v", err)
	}

	for err := range gSrv.GetErrors() {
		log.Fatalf("test server serve err, %v", err)
	}
}
