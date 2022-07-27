package main

import (
	"context"
	"log"

	"banking-app/app"
	"banking-app/app/config"
	"banking-app/app/gateway/http"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfigurations()
	if err != nil {
		log.Fatalf("Could not load environment: %v", err)
	}

	application, err := app.NewApplication(ctx, cfg)
	if err != nil {
		log.Fatalf("Could not set up application: %v", err)
	}

	_ = http.NewServer(application)
	if err != nil {
		log.Fatal("Server shut down")
	}
}
