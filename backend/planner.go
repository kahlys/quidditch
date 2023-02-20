package backend

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
)

type PlannerStore interface {
	EditRecruitablePlayers(context.Context, []Player) error
}

type Planner struct {
	logger *zap.Logger
	store  PlannerStore
}

func NewPlanner(logger *zap.Logger, store PlannerStore) Planner {
	return Planner{
		logger: logger,
		store:  store,
	}
}

func (p Planner) Run() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Hour().Do(p.updateRecruits)
	s.StartImmediately().StartBlocking()

}

func (p Planner) updateRecruits() {
	p.logger.Sugar().Infow("Update recruitable players list")
	if err := p.store.EditRecruitablePlayers(context.TODO(), generatePlayers()); err != nil {
		p.logger.Sugar().Errorw("Update recruitable players list failed", "message", err.Error())
	}
}
