package game

import (
	"log/slog"
	"math/rand"
)

const (
	goalPoints   = 10
	snitchPoints = 150

	roundsMax       = 40000
	roundsMaxSnitch = 1000
)

type Result struct {
	Round     int
	ScoreHome int
	ScoreAway int
	End       bool
}

type Game struct {
	ID       int
	SeasonID int

	Result Result

	Home Team
	Away Team
}

func NewGame(season int, home, away Team) Game {
	return Game{
		SeasonID: season,

		Home: home,
		Away: away,

		Result: Result{Round: 0, ScoreHome: 0, ScoreAway: 0},
	}
}

// Simulate a game between two teams
func (g *Game) Simulate() {
	for failsafe := 0; !g.Result.End || failsafe > 40000; g.Result.Round++ {
		g.simulateRound()
	}

	slog.Info("simulate match", "scoreHome", g.Result.ScoreHome, "scoreAway", g.Result.ScoreAway, "round", g.Result.Round)
}

// simulateRound simulate a round of the game
func (g *Game) simulateRound() {
	g.simulateGoal()
	g.simulateSnitch()
}

// simulateGoal simulate a goal tented by both teams
func (g *Game) simulateGoal() {
	homeChances := g.Home.ChasersPercent()
	awayChances := g.Away.ChasersPercent()

	homeTriesToGoal := rand.Intn(int(homeChances+awayChances)) < int(homeChances)

	switch homeTriesToGoal {
	case true:
		if g.Away.isKeeperStops() {
			g.Result.ScoreHome += goalPoints
		}
	case false:
		if g.Home.isKeeperStops() {
			g.Result.ScoreAway += goalPoints
		}
	}

	slog.Info("simulate goal", "home", g.Home.Name, "scoreHome", g.Result.ScoreHome, "away", g.Away.Name, "scoreAway", g.Result.ScoreAway, "round", g.Result.Round)
}

// simulateSnitch simulate a snitch catch tented by both teams
func (g *Game) simulateSnitch() {
	p := rand.Intn(roundsMaxSnitch)
	if p > g.Result.Round {
		return // snitch not found
	}
	slog.Info("snitch found", "round", g.Result.Round, "snitch", p)

	homeSeekerFirst := g.Home.Squad.Seeker.Power >= g.Away.Squad.Seeker.Power

	hometry := func() {
		if g.seekerFindAndCatchSnitch(g.Home.Squad.Seeker) {
			g.Result.ScoreHome += snitchPoints
			g.Result.End = true
		}
	}
	awaytry := func() {
		if g.seekerFindAndCatchSnitch(g.Away.Squad.Seeker) {
			g.Result.ScoreAway += snitchPoints
			g.Result.End = true
		}
	}

	switch homeSeekerFirst {
	case true:
		hometry()
		if g.Result.End {
			return
		}
		awaytry()
	case false:
		awaytry()
		if g.Result.End {
			return
		}
		hometry()
	}
}

// seekerFindAndCatchSnitch simulate a seeker finding and catching the snitch
func (g *Game) seekerFindAndCatchSnitch(seeker Player) bool {
	return rand.Intn(100) <= seeker.Power // snitch catched
}
