package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kahlys/quidditch/backend"
)

type Database struct {
	*pgxpool.Pool
}

func NewDatabase(connStr string) (*Database, error) {
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return &Database{
		Pool: pool,
	}, nil
}

// User returns user main informations.
func (db *Database) User(userid int) (backend.User, error) {
	user := backend.User{ID: userid}

	tx, err := db.Begin(context.TODO())
	if err != nil {
		return backend.User{}, err
	}
	defer tx.Rollback(context.TODO())

	err = tx.QueryRow(context.TODO(), `SELECT username, email FROM users WHERE id=$1`, user.ID).Scan(&user.Name, &user.Email)
	if err != nil {
		return backend.User{}, err
	}

	err = tx.QueryRow(context.TODO(), `SELECT teams.id FROM teams JOIN users ON teams.owner_id = users.id WHERE users.id = $1`, user.ID).Scan(&user.TeamID)
	if err != nil {
		return backend.User{}, err
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		return backend.User{}, err
	}

	return user, nil
}

// UserByEmail returns user informations and his encrypted password.
func (db *Database) UserByEmail(email string) (backend.User, error) {
	user := backend.User{Email: email}

	tx, err := db.Begin(context.TODO())
	if err != nil {
		return backend.User{}, err
	}
	defer tx.Rollback(context.TODO())

	err = tx.QueryRow(context.TODO(), `SELECT id, username, password FROM users WHERE email=$1`, user.Email).Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return backend.User{}, err
	}

	err = tx.QueryRow(context.TODO(), `SELECT teams.id FROM teams JOIN users ON teams.owner_id = users.id WHERE users.id = $1`, user.ID).Scan(&user.TeamID)
	if err != nil {
		return backend.User{}, err
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		return backend.User{}, err
	}

	return user, nil
}

func (db *Database) UpdateUserLastLogin(userid int) error {
	_, err := db.Exec(context.TODO(), `UPDATE users SET last_login=CURRENT_TIMESTAMP WHERE id=$1`, userid)
	return err
}

// Register registers a user in the database, password must be encrypted.
func (db *Database) RegisterUser(user backend.User, encPassword string, team backend.Team) (int, int, error) {
	userid, teamid := -1, -1

	tx, err := db.Begin(context.TODO())
	if err != nil {
		return userid, teamid, err
	}
	defer tx.Rollback(context.TODO())

	err = tx.QueryRow(
		context.Background(),
		"INSERT INTO users (username, email, password, last_login) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, encPassword, time.Now(),
	).Scan(&userid)
	if err != nil {
		return userid, teamid, err
	}

	err = tx.QueryRow(
		context.Background(),
		`INSERT INTO teams (name, owner_id) VALUES ($1, $2) RETURNING id`,
		team.Name, userid,
	).Scan(&teamid)
	if err != nil {
		return userid, teamid, err
	}

	for _, p := range team.Players() {
		_, err = tx.Exec(
			context.TODO(),
			`INSERT INTO players (first_name, last_name, nationality, power, stamina, position, team_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			p.FirstName, p.LastName, p.Country, p.Power, p.Stamina, p.Role, teamid,
		)
		if err != nil {
			return userid, teamid, err
		}
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		return userid, teamid, err
	}

	return userid, teamid, err
}
