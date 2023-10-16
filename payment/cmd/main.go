package main

import (
	"github.com/kuzminal/microservices/payment/config"
	"github.com/kuzminal/microservices/payment/internal/adapters/db"
	"github.com/kuzminal/microservices/payment/internal/adapters/grpc"
	"github.com/kuzminal/microservices/payment/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
