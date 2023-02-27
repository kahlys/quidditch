package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/kahlys/quidditch/backend"
	"github.com/kahlys/quidditch/backend/store"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := store.NewDatabase("postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	for {
		if err := db.Ping(context.TODO()); err != nil {
			time.Sleep(time.Second)
			continue
		}
		break
	}
	log.Info().Msg("connected to database")

	log.Info().Msg("planner running")
	p := backend.NewPlanner(db)
	p.Run()
}
