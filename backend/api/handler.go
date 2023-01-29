package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	h := handler{}
	return h.handler()
}

type handler struct {
}

func (h handler) handler() http.Handler {
	r := mux.NewRouter()

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

	token, err := generateAccessToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
