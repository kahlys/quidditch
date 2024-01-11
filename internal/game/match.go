package game

import (
	"log/slog"
	"math/rand"
)

type Result struct {
	Round     int
	ScoreHome int
	ScoreAway int
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
	// random score
	g.Result.ScoreHome = rand.Intn(100) * 10
	g.Result.ScoreAway = rand.Intn(100) * 10
	g.Result.Round = rand.Intn(3999)

	slog.Info("simulate match", "home", g.Home.Name, "scoreHome", g.Result.ScoreHome, "away", g.Away.Name, "scoreAway", g.Result.ScoreAway, "round", g.Result.Round)
}
