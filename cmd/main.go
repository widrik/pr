package main

import (
	"log"
	"net"

	"github.com/widrik/pr/internal/config"
	"github.com/widrik/pr/internal/repo"
	grpcserver "github.com/widrik/pr/internal/server"
)

const configFile  = "./config/main.json"

func main() {
	// Config
	configuration, err := config.Init(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Repository
	repo := initRepository(&configuration)

	serversErrorsCh := make(chan error)

	grpcServer := grpcserver.NewServer(&calenderApp, net.JoinHostPort(configuration.GRPCServer.Host, configuration.GRPCServer.Port))
	go func() {
		if err := grpcServer.Start(); err != nil {
			serversErrorsCh <- err
		}
	}()
	defer grpcServer.Stop()

}

func initRepository(configuration *config.Configuration) *repo.Repository {
	r := repo.GetRepo(configuration)
	repo.Migrate(r)

	return r
}