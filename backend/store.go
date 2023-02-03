package backend

type Store interface {
	Register(username, mail, password string) error
}
