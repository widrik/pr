// +build integration

package test

import (
	"log"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/widrik/pr/internal/config"
	"github.com/widrik/pr/internal/repo"
	"github.com/widrik/pr/internal/rotator"
	"github.com/widrik/pr/internal/server"
)

const delay = 5 * time.Second

func TestMain(m *testing.M) {
	log.Printf("wait %s for service will be available...", delay)
	time.Sleep(delay)

	// Проверка конфига
	configuration, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	server := server.New(
		fmt.Sprintf(":%d", configuration.GrpcPort),
		rotator.New(repo.GetRepo(configuration)),
	)

	go server.Start()

	status := m.Run()
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}