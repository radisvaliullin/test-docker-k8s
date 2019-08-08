package testsrv

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/radisvaliullin/test-docker-k8s/proto/api/v1/tsrv"
	"google.golang.org/grpc"
)

// Config keeps server parameters
type Config struct {
	// grpc endpoint
	GAddr string
	// RESTFul endpoint
	HAddr string
}

// Server is grpc server
type Server struct {
	conf Config

	// error chan
	errs chan error

	// grpc
	gsrv *grpc.Server
}

// New init server
func New(c Config) *Server {
	s := &Server{
		conf: c,
		errs: make(chan error, 10),
	}
	return s
}

// Start non blocking server starter
func (s *Server) Start() error {

	// server listener
	ln, err := net.Listen("tcp", s.conf.GAddr)
	if err != nil {
		log.Printf("grpc server listener init, %v", err)
		return err
	}

	// grpc server
	gSrv := grpc.NewServer()
	tsrv.RegisterTestServiceServer(gSrv, s)
	s.gsrv = gSrv
	go func() {
		if err := s.gsrv.Serve(ln); err != nil {
			log.Printf("grpc server serve err, %v", err)
			s.errs <- err
		}
	}()

	// http server
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = tsrv.RegisterTestServiceHandlerFromEndpoint(ctx, mux, s.conf.GAddr, opts)
	if err != nil {
		log.Printf("grcp server, register http endpoint err - %v", err)
		return err
	}

	go func() {
		if err := http.ListenAndServe(s.conf.HAddr, mux); err != nil {
			log.Printf("grpc server, http gateway listen err - %v", err)
			s.errs <- err
		}
	}()

	return nil
}

// GetErrors returns error chan
func (s *Server) GetErrors() <-chan error {
	return s.errs
}
