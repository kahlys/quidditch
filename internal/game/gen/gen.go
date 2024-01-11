package gen

import (
	"math/rand"

	"github.com/kahlys/quidditch/internal/game"
)

// GenerateTeam generate a team with random players
func GenerateTeam(name string, level game.Level) game.Team {
	return game.Team{
		Name: name,
		Squad: game.Squad{
			Seeker:  generatePlayerWithRole(level, game.RoleSeeker),
			Keeper:  generatePlayerWithRole(level, game.RoleKeeper),
			Beater1: generatePlayerWithRole(level, game.RoleBeater),
			Beater2: generatePlayerWithRole(level, game.RoleBeater),
			Chaser1: generatePlayerWithRole(level, game.RoleChaser),
			Chaser2: generatePlayerWithRole(level, game.RoleChaser),
			Chaser3: generatePlayerWithRole(level, game.RoleChaser),
		},
	}
}

// GeneratePlayers generate n players with random stats according to their level
func GeneratePlayers(n int, level game.Level) []game.Player {
	players := []game.Player{}
	for i := 0; i < n; i++ {
		players = append(players, generatePlayer(level))
	}
	return players
}

// generatePlayer generate a player with random stats according to his level
func generatePlayer(level game.Level) game.Player {
	country := Countries[rand.Intn(len(Countries))]
	p := game.Player{
		FirstName: firstNameByCountry[country][rand.Intn(len(firstNameByCountry[country]))],
		LastName:  lastNameByCountry[country][rand.Intn(len(lastNameByCountry[country]))],
		Role:      game.Roles[rand.Intn(len(game.Roles))],
		Country:   country,
		Power:     generateStat(level),
		Stamina:   100,
	}

	return p
}

// generatePlayerWithRole generate a player with a specific role
func generatePlayerWithRole(level game.Level, role string) game.Player {
	p := generatePlayer(level)
	p.Role = role
	return p
}

func generateStat(level game.Level) int {
	switch level {
	case game.Level1:
		return randInt(20, 40)
	case game.Level2:
		return randInt(20, 50)
	case game.Level3:
		return randInt(20, 60)
	case game.Level4:
		return randInt(20, 70)
	case game.Level5:
		return randInt(20, 80)
	default:
		return randInt(10, 90)
	}
}

func randInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
