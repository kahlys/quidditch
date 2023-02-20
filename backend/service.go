package backend

import (
	"context"
	"encoding/hex"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Store interface {
	UpdateUserLastLogin(userid int) error
	User(userid int) (User, error)
	UserByEmail(email string) (User, error)
	RegisterUser(user User, encPassword string, team Team) (userID int, teamID int, err error)
	Team(teamid int) (Team, error)

	RecruitablePlayers(context.Context) ([]Player, error)
}

type Service struct {
	logger *zap.Logger

	Store Store
}

func NewService(logger *zap.Logger, store Store) *Service {
	return &Service{
		logger: logger,
		Store:  store,
	}
}

// Register an user and return it's id.
func (s *Service) CreateUser(user User) (int, int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, -1, err
	}

	team := GenerateFirstTeam(fmt.Sprintf("%v team", user.Name))

	userid, teamid, err := s.Store.RegisterUser(user, hex.EncodeToString(hashedPassword), team)
	if err != nil {
		return -1, -1, err
	}

	return userid, teamid, err
}

func (s *Service) AuthUser(user User) (User, error) {
	u, err := s.Store.UserByEmail(user.Email)
	if err != nil {
		return User{}, err
	}

	pass, err := hex.DecodeString(u.Password)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword(pass, []byte(user.Password))
	if err != nil {
		return User{}, err
	}

	err = s.Store.UpdateUserLastLogin(u.ID)
	if err != nil {
		s.logger.Sugar().Warnln("unable to update last_login")
	}

	return u, nil
}

// Team return the team name and players.
func (s *Service) Team(teamid int) (Team, error) {
	return s.Store.Team(teamid)
}

// RecruitablePlayers return available players to recruit.
func (s *Service) RecruitablePlayers() ([]Player, error) {
	return s.Store.RecruitablePlayers(context.TODO())
}
