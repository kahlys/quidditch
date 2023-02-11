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
	var user backend.User

	// Decode the incoming JSON data into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userid, teamid, err := h.s.CreateUser(user)
	if err != nil {
		h.logger.Sugar().Errorw("user registration", "message", err.Error())
		http.Error(w, "user registration", http.StatusInternalServerError)
		return
	}

	token, err := generateAccessToken(userid, teamid)
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
	var user backend.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = h.s.AuthUser(user)
	if err != nil {
		h.logger.Sugar().Errorw("user login", "message", err.Error())
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := generateAccessToken(user.ID, user.TeamID)
	if err != nil {
		http.Error(w, "token generation", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
}

// home handles home endpoint
func (h handler) home(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(ctxUserID).(int)
	teamID := r.Context().Value(ctxTeamID).(int)

	w.Write([]byte(fmt.Sprintf("Welcome user %v with team %v", userID, teamID)))
}
