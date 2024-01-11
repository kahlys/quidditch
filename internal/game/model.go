package game

const (
	RoleSeeker = "seeker"
	RoleChaser = "chaser"
	RoleBeater = "beater"
	RoleKeeper = "keeper"
)

var Roles = []string{
	RoleSeeker,
	RoleChaser,
	RoleBeater,
	RoleKeeper,
}

type Level int

const (
	Level1 Level = iota + 1
	Level2
	Level3
	Level4
	Level5
)

// User struct to hold user information
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TeamID   int    `json:"-"`
}

// Player describe a player with his infos and stats.
type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	Role      string `json:"role"`
	Stamina   int    `json:"stamina"`
	Power     int    `json:"power"`
}

type Squad struct {
	Seeker  Player
	Chaser1 Player
	Chaser2 Player
	Chaser3 Player
	Beater1 Player
	Beater2 Player
	Keeper  Player
}

type Team struct {
	ID    int
	Name  string
	Squad Squad
}

func (t Team) Players() []Player {
	return []Player{
		t.Squad.Seeker,
		t.Squad.Keeper,
		t.Squad.Beater1,
		t.Squad.Beater2,
		t.Squad.Chaser1,
		t.Squad.Chaser2,
		t.Squad.Chaser3,
	}
}
