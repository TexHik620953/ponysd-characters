package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"ponysd-characters/appconfig"
	"ponysd-characters/application"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := appconfig.LoadAppConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := application.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to init application: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
