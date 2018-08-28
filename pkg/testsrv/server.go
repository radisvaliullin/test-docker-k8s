package testsrv

import "google.golang.org/grpc"

// Config keeps server parameters
type Config struct {
	Addr string
}

// Server is grpc server
type Server struct {
	conf Config

	// grpc
	gsrv *grpc.Server
}

// New init server
func New(c Config) *Server {
	s := &Server{
		conf: c,
	}
	return s
}

// Start non blocking server starter
func (s *Server) Start() error {

	gSrv := grpc.NewServer()
	s.gsrv = gSrv
}
