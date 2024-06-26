package main

import (
	"context"
	"log/slog"

	"exercise3/problem1/internal/api"
	"exercise3/problem1/internal/conf"
	"exercise3/problem1/internal/db"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// conf
	configs, err := conf.NewConfig(ctx)
	if err != nil {
		panic(err)
	}

	// configure db service
	m, err := db.NewModel(configs.DB)
	if err != nil {
		slog.Error("Failed to connect to database", err)
	}
	slog.InfoContext(ctx, "initialize service", "service", "db")

	// configure gateway service
	srv := api.NewApi(configs.Api, m)
	slog.InfoContext(ctx, "initialize service", "service", "api")
	// start gateway service
	srv.Start(ctx, cancel)

	<-ctx.Done()
	// Your cleanup tasks go here
	slog.InfoContext(ctx, "cleaning up ...")

	slog.InfoContext(ctx, "server was successful shutdown.")
}
