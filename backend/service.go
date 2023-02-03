package backend

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Store Store
}

func (s *Service) Register(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.Store.Register(username, email, hex.EncodeToString(hashedPassword))
}

func (s *Service) Login(email, password string) error {
	return nil
}
