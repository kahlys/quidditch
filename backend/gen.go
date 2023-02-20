package backend

import "math/rand"

func GenerateFirstTeam(name string) Team {
	return Team{
		Name: name,
		Squad: Squad{
			Seeker:  generatePlayerWithRole(1, RoleSeeker),
			Keeper:  generatePlayerWithRole(1, RoleKeeper),
			Beater1: generatePlayerWithRole(1, RoleBeater),
			Beater2: generatePlayerWithRole(1, RoleBeater),
			Chaser1: generatePlayerWithRole(1, RoleChaser),
			Chaser2: generatePlayerWithRole(1, RoleChaser),
			Chaser3: generatePlayerWithRole(1, RoleChaser),
		},
	}
}

func generatePlayers() []Player {
	players := []Player{}
	for level := 0; level <= 5; level++ {
		for i := 0; i < 10; i++ {
			players = append(players, generatePlayer(level))
		}
	}
	return players
}

func generatePlayer(level int) Player {
	country := countries[rand.Intn(len(countries))]
	p := Player{
		FirstName: firstNameByCountry[country][rand.Intn(len(firstNameByCountry[country]))],
		LastName:  lastNameByCountry[country][rand.Intn(len(lastNameByCountry[country]))],
		Role:      roles[rand.Intn(len(roles))],
		Country:   country,
		Power:     generateStat(level),
		Stamina:   100,
	}

	return p
}

func generatePlayerWithRole(level int, role string) Player {
	p := generatePlayer(level)
	p.Role = role
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
