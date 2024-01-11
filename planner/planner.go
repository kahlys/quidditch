package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-co-op/gocron/v2"

	"github.com/kahlys/quidditch/internal/game"
	"github.com/kahlys/quidditch/internal/game/gen"
	"github.com/kahlys/quidditch/internal/store"
)

type Config struct {
	Store store.Database
}

type Planner struct {
	store store.Database

	stop chan struct{} // use to stop the planner
}

func NewPlanner(cfg Config) Planner {
	return Planner{
		stop:  make(chan struct{}),
		store: cfg.Store,
	}
}

func (p *Planner) Run() error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	_, err = s.NewJob(
		gocron.DurationJob(1*time.Hour),
		gocron.NewTask(p.updateRecruits),
		gocron.WithStartAt(gocron.WithStartDateTime(nextHour())),
	)
	if err != nil {
		return err
	}

	_, err = s.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(15, 00, 00))),
		gocron.NewTask(p.runMatches),
	)
	if err != nil {
		return err
	}

	s.Start()
	<-p.stop

	return nil
}

func nextHour() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())
}

// Stop the planner
func (p *Planner) Stop() {
	close(p.stop)
}

// updateRecruits update the list of recruitable players
func (p *Planner) updateRecruits() {
	slog.Info("update recruitable players list")

	players := []game.Player{}

	players = append(players, gen.GeneratePlayers(10, game.Level1)...)
	players = append(players, gen.GeneratePlayers(10, game.Level2)...)
	players = append(players, gen.GeneratePlayers(5, game.Level3)...)
	players = append(players, gen.GeneratePlayers(4, game.Level4)...)
	players = append(players, gen.GeneratePlayers(1, game.Level5)...)

	if err := p.store.NewRecruitablePlayers(context.Background(), players); err != nil {
		slog.Error("update recruitable players list failed", "error", err)
	}
}

func (p *Planner) runMatches() {
	matches, err := p.store.Matches(context.Background())
	if err != nil {
		slog.Error("unable to get matches", "error", err)
		return
	}

	for _, m := range matches {
		match := m
		go func() {
			match.Simulate()
			err := p.store.SaveMatchResult(context.Background(), match)
			if err != nil {
				slog.Error("unable to update match", "error", err)
			}
		}()
	}
}
