package main

import (
	"WorkoutTracker/internal/database"
	"errors"
	"fmt"
	"github.com/alexedwards/argon2id"
	"log"
	"net/http"
	"runtime"
)

func hashing(password string) string {
	params := &argon2id.Params{
		Memory:      128 * 1024,
		Iterations:  4,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		KeyLength:   32,
	}
	hash, erro := argon2id.CreateHash(password, params)
	if erro != nil {
		log.Fatal(erro)
	}
	return hash
}

func (app *application) Session(w http.ResponseWriter, r *http.Request) (sess database.Sessions, err error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// errors.Is() checks if error matches a specific type even if its wrapped inside another error
		if errors.Is(err, http.ErrNoCookie) {
			return database.Sessions{}, errors.New("No session token cookie")
		}
		return database.Sessions{}, fmt.Errorf("Error getting session token: %w", err)
	}
	s, err := app.Models.UserModel.QUERYSESSION(cookie.Value)
	if err != nil {
		return database.Sessions{}, errors.New("Invalid or expired Session")
	}
	return *s, nil
}
