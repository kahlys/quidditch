package game

import (
	"math/rand"
	"testing"

	"go.uber.org/zap"

	"github.com/kahlys/quidditch/backend"
	"github.com/stretchr/testify/require"
)

func TestGame(t *testing.T) {
	rand.Seed(0)

	teamHome := backend.Team{
		Name: "Home",
		Squad: backend.Squad{
			Seeker:  backend.Player{FirstName: "home-seeker", Power: 10, Stamina: 100},
			Chaser1: backend.Player{FirstName: "home-chaser-1", Power: 10, Stamina: 100},
			Chaser2: backend.Player{FirstName: "home-chaser-2", Power: 10, Stamina: 100},
			Chaser3: backend.Player{FirstName: "home-chaser-3", Power: 10, Stamina: 100},
			Beater1: backend.Player{FirstName: "home-beater-1", Power: 10, Stamina: 100},
			Beater2: backend.Player{FirstName: "home-beater-2", Power: 10, Stamina: 100},
			Keeper:  backend.Player{FirstName: "home-keeper", Power: 10, Stamina: 100},
		},
	}

	teamAway := backend.Team{
		Name: "Away",
		Squad: backend.Squad{
			Seeker:  backend.Player{FirstName: "away-seeker", Power: 90, Stamina: 100},
			Chaser1: backend.Player{FirstName: "away-chaser-1", Power: 90, Stamina: 100},
			Chaser2: backend.Player{FirstName: "away-chaser-2", Power: 90, Stamina: 100},
			Chaser3: backend.Player{FirstName: "away-chaser-3", Power: 90, Stamina: 100},
			Beater1: backend.Player{FirstName: "away-beater-1", Power: 90, Stamina: 100},
			Beater2: backend.Player{FirstName: "away-beater-2", Power: 90, Stamina: 100},
			Keeper:  backend.Player{FirstName: "away-keeper", Power: 90, Stamina: 100},
		},
	}

	game := NewGame(zap.NewExample(), 1, teamHome, teamAway)
	game.Simulate()

	results := game.Results()
	require.Greater(t, results.ScoreAway, results.ScoreHome)
}
