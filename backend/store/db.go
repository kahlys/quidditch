package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
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

// Register registers a user in the database, password must be encrypted.
func (db *Database) Register(username, email, password string) error {
	_, err := db.Exec(
		context.Background(),
		"INSERT INTO users (username, email, password, last_login) VALUES ($1, $2, $3, $4)",
		username, email, password, time.Now(),
	)
	return err
}
