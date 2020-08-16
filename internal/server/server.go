package server

import (
	"net"

	"github.com/widrik/pr/api/spec"
	"github.com/widrik/pr/internal/rotator"
	"google.golang.org/grpc"
)

type Server struct {
	address       string
	grpcServer    *grpc.Server
}

func New(address string, r *rotator.Rotator) *Server {
	srv := &Server{
		address: address,
	}

	rotatorClient := rotator.NewRotatorClient(r)

	grpcServer := grpc.NewServer()
	spec.RegisterRotationServiceServer(grpcServer, rotatorClient)
	srv.grpcServer = grpcServer

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
