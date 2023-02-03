package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/kahlys/quidditch/backend"
)

func Handler(logger *zap.Logger, s *backend.Service) http.Handler {
	h := handler{
		logger: logger,
		s:      s,
	}
	return h.handler()
}

type handler struct {
	logger *zap.Logger

	s *backend.Service
}

func (h handler) handler() http.Handler {
	r := mux.NewRouter().PathPrefix("/api/").Subrouter()

	r.HandleFunc("/register", h.register).Methods("POST")
	r.HandleFunc("/login", h.login).Methods("POST")

	rAuth := r.NewRoute().Subrouter()
	rAuth.Use(h.mwAuth)
	rAuth.HandleFunc("/home", h.home).Methods("GET")

	return r
}

// register handles the registration of new users
func (h handler) register(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode the incoming JSON data into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: save in db (and get id)
	user.ID = len(Users) + 1
	Users = append(Users, user)

	err = h.s.Register(user.Name, user.Email, user.Password)
	if err != nil {
		h.logger.Sugar().Errorw("user registration", "message", err.Error())
		http.Error(w, "user registration", http.StatusInternalServerError)
		return
	}

	token, err := generateAccessToken(user.ID)
	if err != nil {
		h.logger.Sugar().Errorw("token generation", "message", err.Error())
		http.Error(w, "token generation", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
}

// login handles the login of existing users
func (h handler) login(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: get from db
	for _, u := range Users {
		if u.Email == user.Email && u.Password == user.Password {
			token, err := generateAccessToken(u.ID)
			if err != nil {
				http.Error(w, "token generation", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "token",
				Value: token,
			})
			return
		}
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

// home handles home endpoint
func (h handler) home(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ctxUserID).(int)

	w.Write([]byte(fmt.Sprintf("Welcome user %v", userID)))
}
