package main

import (
	"log"

	"github.com/radisvaliullin/test-docker-k8s/pkg/testsrv"
)

var addr = "0.0.0.0:7373"

func main() {

	log.Printf("test server start, addr - %v", addr)

	conf := testsrv.Config{Addr: addr}
	gSrv := testsrv.New(conf)
	err := gSrv.Start()
	if err != nil {
		log.Fatalf("test server start err, %v", err)
	}

	for err := range gSrv.GetErrors() {
		log.Fatalf("test server serve err, %v", err)
	}
}
