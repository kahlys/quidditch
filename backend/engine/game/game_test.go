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
			Seeker:  engine.Player{Name: "home-seeker", Power: 10, Stamina: 100},
			Chaser1: engine.Player{Name: "home-chaser-1", Power: 10, Stamina: 100},
			Chaser2: engine.Player{Name: "home-chaser-2", Power: 10, Stamina: 100},
			Chaser3: engine.Player{Name: "home-chaser-3", Power: 10, Stamina: 100},
			Beater1: engine.Player{Name: "home-beater-1", Power: 10, Stamina: 100},
			Beater2: engine.Player{Name: "home-beater-2", Power: 10, Stamina: 100},
			Keeper:  engine.Player{Name: "home-keeper", Power: 10, Stamina: 100},
		},
	}

	teamAway := engine.Team{
		Name: "Away",
		Squad: engine.Squad{
			Seeker:  engine.Player{Name: "away-seeker", Power: 90, Stamina: 100},
			Chaser1: engine.Player{Name: "away-chaser-1", Power: 90, Stamina: 100},
			Chaser2: engine.Player{Name: "away-chaser-2", Power: 90, Stamina: 100},
			Chaser3: engine.Player{Name: "away-chaser-3", Power: 90, Stamina: 100},
			Beater1: engine.Player{Name: "away-beater-1", Power: 90, Stamina: 100},
			Beater2: engine.Player{Name: "away-beater-2", Power: 90, Stamina: 100},
			Keeper:  engine.Player{Name: "away-keeper", Power: 90, Stamina: 100},
		},
	}

	game := NewGame(zap.NewExample(), 1, teamHome, teamAway)
	game.Simulate()

	results := game.Results()
	require.Greater(t, results.ScoreAway, results.ScoreHome)
}
