package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/kahlys/quidditch/internal/store"
)

func main() {
	db, err := store.NewDatabase("postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	time.Sleep(10 * time.Second)

	cfg := Config{
		Store: db,
	}

	slog.Info("starting planner")
	p := NewPlanner(cfg)
	p.Run()
}
