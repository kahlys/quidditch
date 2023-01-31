package game

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/kahlys/quidditch/backend/engine"
)

func TestGame(t *testing.T) {
	rand.Seed(0)

	teamHome := engine.Team{
		Name: "Home",
		Squad: engine.Squad{
			Seeker:  engine.Player{FirstName: "home-seeker", Power: 10, Stamina: 100},
			Chaser1: engine.Player{FirstName: "home-chaser-1", Power: 10, Stamina: 100},
			Chaser2: engine.Player{FirstName: "home-chaser-2", Power: 10, Stamina: 100},
			Chaser3: engine.Player{FirstName: "home-chaser-3", Power: 10, Stamina: 100},
			Beater1: engine.Player{FirstName: "home-beater-1", Power: 10, Stamina: 100},
			Beater2: engine.Player{FirstName: "home-beater-2", Power: 10, Stamina: 100},
			Keeper:  engine.Player{FirstName: "home-keeper", Power: 10, Stamina: 100},
		},
	}

	teamAway := engine.Team{
		Name: "Away",
		Squad: engine.Squad{
			Seeker:  engine.Player{FirstName: "away-seeker", Power: 90, Stamina: 100},
			Chaser1: engine.Player{FirstName: "away-chaser-1", Power: 90, Stamina: 100},
			Chaser2: engine.Player{FirstName: "away-chaser-2", Power: 90, Stamina: 100},
			Chaser3: engine.Player{FirstName: "away-chaser-3", Power: 90, Stamina: 100},
			Beater1: engine.Player{FirstName: "away-beater-1", Power: 90, Stamina: 100},
			Beater2: engine.Player{FirstName: "away-beater-2", Power: 90, Stamina: 100},
			Keeper:  engine.Player{FirstName: "away-keeper", Power: 90, Stamina: 100},
		},
	}

	game := NewGame(zap.NewExample(), 1, teamHome, teamAway)
	game.Simulate()

	results := game.Results()
	require.Greater(t, results.ScoreAway, results.ScoreHome)
}
