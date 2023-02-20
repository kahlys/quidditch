package main

import (
	"context"
	"os"
	"time"

	"github.com/kahlys/quidditch/backend"
	"github.com/kahlys/quidditch/backend/store"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		os.Exit(1)
	}

	db, err := store.NewDatabase("postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		logger.Sugar().Fatalw("failed to connect to database", "message", err)
	}

	for {
		if err := db.Ping(context.TODO()); err != nil {
			time.Sleep(time.Second)
			continue
		}
		break
	}
	logger.Sugar().Infow("connected to database")

	logger.Sugar().Infow("planner running")
	p := backend.NewPlanner(logger, db)
	p.Run()
}
