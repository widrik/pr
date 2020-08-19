package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/widrik/pr/internal/config"
	"github.com/widrik/pr/internal/repo"
	"github.com/widrik/pr/internal/rotator"
	"github.com/widrik/pr/internal/server"
)

func main() {
	configuration, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	r := rotator.New(initRepository(configuration))

	serversErrorsCh := make(chan error)
	grpcServer := server.New(fmt.Sprintf(":%d", configuration.GrpcPort), r)
	go func() {
		if err := grpcServer.Start(); err != nil {
			serversErrorsCh <- err
		}
	}()
	defer grpcServer.Stop()

	print("Working...")
	signalsCh := make(chan os.Signal, 1)
	signal.Notify(signalsCh, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalsCh:
		signal.Stop(signalsCh)

		return
	case err = <-serversErrorsCh:
		if err != nil {
			log.Fatal(err)
		}

		return
	}
}

func initRepository(configuration *config.Configuration) *repo.Repository {
	r := repo.GetRepo(configuration)
	repo.Migrate(r)

	return r
}
