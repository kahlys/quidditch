package backend

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

type PlannerStore interface {
	NewRecruitablePlayers(context.Context, []Player) error
	InitBotTeam(context.Context, Team) error
	Teams(context.Context) ([]Team, error)
	Matches(ctx context.Context, seaon, day int) ([]*Game, error)
}

type Planner struct {
	store PlannerStore

	season int
	day    int
}

func NewPlanner(store PlannerStore) *Planner {
	return &Planner{
		store: store,
	}
}

func (p *Planner) Run() {
	p.Init()

	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Hour().Do(p.updateRecruits)

	s.Every(1).Minutes().Do(p.matches)

	s.StartImmediately().StartBlocking()
}

func (p *Planner) Init() {
	p.season = 0 // TODO: get from db
	p.day = 0    // TODO: get from db

	// TOREMOVE: init bot teams
	teams, err := p.store.Teams(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Msg("unable to get teams")
	}
	if len(teams) == 0 {
		for _, n := range []string{"Gryffindor", "Hufflepuff", "Ravenclaw", "Slytherin"} {
			err = p.store.InitBotTeam(context.TODO(), GenerateFirstTeam(n))
			if err != nil {
				log.Fatal().Str("name", n).Err(err).Msg("unable to save bot team")
			}
		}
	}
}

func (p *Planner) updateRecruits() {
	log.Info().Msg("Update recruitable players list")
	if err := p.store.NewRecruitablePlayers(context.TODO(), generatePlayers()); err != nil {
		log.Error().Err(err).Msg("Update recruitable players list failed")
	}
}

var ErrRestDay = fmt.Errorf("rest day")

func (p *Planner) matches() {
	p.day++
	log.Info().Int("season", p.season).Int("day", p.day).Msg("Simulate matches")

	games, err := p.store.Matches(context.TODO(), p.season, p.day)
	if errors.Is(err, ErrRestDay) {
		p.matches()
		return
	} else if err != nil {
		log.Error().Int("season", p.season).Int("day", p.day).Err(err).Msg("Simulate matches failed")
	}

	for _, g := range games {
		g.Simulate()
		fmt.Println(g.Results())
	}
	fmt.Println(games)
}
