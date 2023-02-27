package backend

import (
	"math/rand"

	"github.com/rs/zerolog/log"
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
	id     int
	result Result

	home Team
	away Team
}

func NewGame(id int, home, away Team) *Game {
	return &Game{
		home: home,
		away: away,

		result: Result{End: false, Round: 0, ScoreHome: 0, ScoreAway: 0},
	}
}

// Results return the game results
func (g *Game) Results() Result { return g.result }

// Simulate a game between two teams.
func (g *Game) Simulate() {
	log.Info().Str("home", g.home.Name).Str("away", g.away.Name).Int("id", g.id).Msg("Game started")

	failsafe := 0
	for ; !g.result.End; g.result.Round++ {
		g.simulateRound()
		failsafe++
		if failsafe > 40000 {
			log.Warn().Msg("game has been stopped timeout")
			break
		}
	}

	log.Info().Str("home", g.home.Name).Str("away", g.away.Name).Int("id", g.id).Msg("Game ended")
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

	g.result.End = g.simulateRoundSnitch()
}

// simulateRoundSnitch simulates when seekers try to find and catch the golden snitch. Return true if the snitch has been caught.
func (g *Game) simulateRoundSnitch() bool {
	if g.home.Squad.Seeker.Power >= g.away.Squad.Seeker.Power {
		if seekerFindAndCatchSnitch(g.home.Squad.Seeker) {
			log.Debug().Str("seeker", g.home.Squad.Seeker.FirstName).Int("id", g.id).Msg("Seeker catch the golden snitch")
			g.result.ScoreHome += scoreSnitch
			return true
		} else if seekerFindAndCatchSnitch(g.away.Squad.Seeker) {
			log.Debug().Str("seeker", g.away.Squad.Seeker.FirstName).Int("id", g.id).Msg("Seeker catch the golden snitch")
			g.result.ScoreAway += scoreSnitch
			return true
		}
	} else {
		if seekerFindAndCatchSnitch(g.away.Squad.Seeker) {
			log.Debug().Str("seeker", g.away.Squad.Seeker.FirstName).Int("id", g.id).Msg("Seeker catch the golden snitch")
			g.result.ScoreAway += scoreSnitch
			return true
		} else if seekerFindAndCatchSnitch(g.home.Squad.Seeker) {
			log.Debug().Str("seeker", g.home.Squad.Seeker.FirstName).Int("id", g.id).Msg("Seeker catch the golden snitch")
			g.result.ScoreHome += scoreSnitch
			return true
		}
	}
	return false
}

func snitchAppears(chance int) bool {
	return diceRoll(chance, 100, 1)
}

func seekerFindAndCatchSnitch(player Player) bool {
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
