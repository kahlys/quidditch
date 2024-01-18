package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame(t *testing.T) {
	teamHome := Team{
		Name: "Home",
		Squad: Squad{
			Seeker:  Player{FirstName: "home-seeker", Power: 90, Stamina: 100},
			Chaser1: Player{FirstName: "home-chaser-1", Power: 90, Stamina: 100},
			Chaser2: Player{FirstName: "home-chaser-2", Power: 90, Stamina: 100},
			Chaser3: Player{FirstName: "home-chaser-3", Power: 90, Stamina: 100},
			Beater1: Player{FirstName: "home-beater-1", Power: 90, Stamina: 100},
			Beater2: Player{FirstName: "home-beater-2", Power: 90, Stamina: 100},
			Keeper:  Player{FirstName: "home-keeper", Power: 90, Stamina: 100},
		},
	}

	teamAway := Team{
		Name: "Away",
		Squad: Squad{
			Seeker:  Player{FirstName: "away-seeker", Power: 90, Stamina: 100},
			Chaser1: Player{FirstName: "away-chaser-1", Power: 90, Stamina: 100},
			Chaser2: Player{FirstName: "away-chaser-2", Power: 90, Stamina: 100},
			Chaser3: Player{FirstName: "away-chaser-3", Power: 90, Stamina: 100},
			Beater1: Player{FirstName: "away-beater-1", Power: 90, Stamina: 100},
			Beater2: Player{FirstName: "away-beater-2", Power: 90, Stamina: 100},
			Keeper:  Player{FirstName: "away-keeper", Power: 90, Stamina: 100},
		},
	}

	game := NewGame(1, teamHome, teamAway)
	game.Simulate()

	results := game.Result
	require.Greater(t, results.ScoreAway+ results.ScoreHome,snitchPoints)
}