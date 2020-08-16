package server

import (
	"github.com/widrik/pr/api/spec"
	"net"

	"github.com/widrik/pr/internal/rotator"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer	  *grpc.Server
	rotatorServer *rotator.Server
	address       string
}

func NewServer(r *rotator.Rotator, listenAddress string) *Server {
	grpcServer := grpc.NewServer()

	srv := &Server{}

	srv.rotatorServer = rotator.NewServer(r)
	srv.grpcServer = grpcServer
	srv.address = listenAddress

	spec.RegisterRotationServiceServer(grpcServer, srv.rotatorServer)

	return srv
}

func (srv *Server) Start() error {
	lis, err := net.Listen("tcp", srv.address)
	if err != nil {
		return err
	}
	err = srv.grpcServer.Serve(lis)

	return err
}

func (srv *Server) Stop() {
	srv.grpcServer.GracefulStop()
}