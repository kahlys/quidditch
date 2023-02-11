package backend

type Store interface {
	UpdateUserLastLogin(userid int) error
	User(userid int) (User, error)
	UserByEmail(email string) (User, error)
	RegisterUser(user User, encPassword string, team Team) (userID int, teamID int, err error)
	Team(teamid int) (Team, error)
}
