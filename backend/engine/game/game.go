package game

import (
	"math/rand"

	"go.uber.org/zap"

	"github.com/kahlys/quidditch/backend/engine"
)

const (
	// number of tries you need to succeed
	triesSnitch = 6

	scoreSnitch = 150
	scoreGoal   = 10
)

type Result struct {
	End       bool
	Round     int
	ScoreHome int
	ScoreAway int
}

type Game struct {
	logger *zap.Logger

	id     int
	result Result

	home engine.Team
	away engine.Team
}

func NewGame(logger *zap.Logger, id int, home, away engine.Team) *Game {
	return &Game{
		logger: logger,

		id: id,

		home: home,
		away: away,

		result: Result{End: false, Round: 0, ScoreHome: 0, ScoreAway: 0},
	}
}

// Results return the game results
func (g *Game) Results() Result { return g.result }

// Simulate a game between two teams.
func (g *Game) Simulate() {
	g.logger.Sugar().Infow("Game started", "gameID", g.id, "home", g.home.Name, "away", g.away.Name)

	failsafe := 0
	for ; !g.result.End; g.result.Round++ {
		g.simulateRound()
		failsafe++
		if failsafe > 40 {
			break
		}
	}

	g.logger.Sugar().Infow("Game ended", "gameID", g.id, "home", g.home.Name, "away", g.away.Name, "scoreHome", g.result.ScoreHome, "scoreAway", g.result.ScoreAway)
}

// simulateRound simulate a try when chasers of a team try to goal.
func (g *Game) simulateRound() {
	// chasers of the current team try to do 2 pass and goal
	// TODO

	// beaters try to take down an opponent
	// TODO

	// seekers try to catch the snitch
	if !snitchAppears(g.result.Round) {
		return
	}
	g.logger.Sugar().Debugw("Snitch Appears", "gameID", g.id, "snitchChance", g.result.Round)
	g.result.End = g.simulateRoundSnitch()
}

// simulateRoundSnitch simulates when seekers try to find and catch the golden snitch. Return true if the snitch has been caught.
func (g *Game) simulateRoundSnitch() bool {
	if g.home.Squad.Seeker.Power >= g.away.Squad.Seeker.Power {
		if seekerFindAndCatchSnitch(g.home.Squad.Seeker) {
			g.logger.Sugar().Debugw("Seeker catch the golden snitch", "gameID", g.id, "team", g.home.Name, "seeker", g.home.Squad.Seeker.Name)
			g.result.ScoreHome += scoreSnitch
			return true
		} else if seekerFindAndCatchSnitch(g.away.Squad.Seeker) {
			g.logger.Sugar().Debugw("Seeker catch the golden snitch", "gameID", g.id, "team", g.away.Name, "seeker", g.away.Squad.Seeker.Name)
			g.result.ScoreAway += scoreSnitch
			return true
		}
	} else {
		if seekerFindAndCatchSnitch(g.away.Squad.Seeker) {
			g.logger.Sugar().Debugw("Seeker catch the golden snitch", "gameID", g.id, "team", g.away.Name, "seeker", g.away.Squad.Seeker.Name)
			g.result.ScoreAway += scoreSnitch
			return true
		} else if seekerFindAndCatchSnitch(g.home.Squad.Seeker) {
			g.logger.Sugar().Debugw("Seeker catch the golden snitch", "gameID", g.id, "team", g.home.Name, "seeker", g.home.Squad.Seeker.Name)
			g.result.ScoreHome += scoreSnitch
			return true
		}
	}
	return false
}

func snitchAppears(chance int) bool {
	return diceRoll(chance, 100, 1)
}

func seekerFindAndCatchSnitch(player engine.Player) bool {
	return diceRoll(player.Power, player.Stamina, triesSnitch)
}

func diceRoll(stat int, stamina int, nb int) bool {
	for i := 0; i < nb; i++ {
		if rand.Intn(101) > stat-(100-stamina) {
			return false
		}
	}
	return true
}
