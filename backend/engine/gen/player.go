package gen

import (
	"math/rand"

	"github.com/kahlys/quidditch/backend/engine"
)

func generatePlayers() []engine.Player {
	players := []engine.Player{}
	for i := 0; i < 10; i++ {
		players = append(players, generatePlayer(0))
	}
	return players
}

func generatePlayer(level int) engine.Player {
	country := countries[rand.Intn(len(countries))]
	p := engine.Player{
		FirstName: firstNameByCountry[country][rand.Intn(len(firstNameByCountry[country]))],
		LastName:  lastNameByCountry[country][rand.Intn(len(lastNameByCountry[country]))],
		Role:      roles[rand.Intn(len(roles))],
		Country:   country,
		Power:     generateStat(level),
		Stamina:   100,
	}

	return p
}

func generateStat(level int) int {
	switch level {
	case 1:
		return randInt(20, 40)
	case 2:
		return randInt(30, 50)
	case 3:
		return randInt(40, 60)
	case 4:
		return randInt(50, 70)
	case 5:
		return randInt(60, 80)
	default:
		return randInt(10, 90)
	}
}

func randInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
