package main

import (
	internal "WorkoutTracker/internal/database"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
	"log"
	"net/http"
	"runtime"
)

type PasswordHasher interface {
	Compare(password, hash string) (bool, error)
}

type Argon2Hasher struct{}

func (a Argon2Hasher) Compare(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}

// DefaultPasswordHasher implements PasswordHasher using argon2id
type DefaultPasswordHasher struct{}

// ComparePasswordAndHash calls argon2id.ComparePasswordAndHash
func (d DefaultPasswordHasher) ComparePasswordAndHash(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}

func Hashing(password string) string {
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

func (app *application) Session(w http.ResponseWriter, r *http.Request) (sess internal.Sessions, err error) {
	cookie, err := r.Cookie("session_token")
	if cookie == nil || err != nil {
		fmt.Println("cookie not found/error occurred while getting session_token")
	} else {
		fmt.Println("cookie found " + cookie.Value)
	}
	if err != nil {
		// errors.Is() checks if error matches a specific type even if its wrapped inside another error
		if errors.Is(err, http.ErrNoCookie) {
			return internal.Sessions{}, errors.New("No session token cookie")
		}
		return internal.Sessions{}, fmt.Errorf("Error getting session token: %w", err)
	}
	s, err := app.Models.UserModel.QuerySession(cookie.Value)
	if err != nil {
		return internal.Sessions{}, errors.New("Invalid or expired Session")
	}
	return *s, nil
}

// helper function for JSON error responses
// tells the client(browser,frontend,API consumer) that the response is in JSON
// without this the client might assume its plain text
func (app *application) sendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (app *application) sendSuccessResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func MigrateDB(db *gorm.DB) {
	//creating table using automigrate
	err := db.AutoMigrate(&internal.Users{}, &internal.Sessions{}, &internal.Exercises{}, &internal.WorkoutToUser{}, &internal.Workouts{}, &internal.UserBests{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}
	log.Println("Database migrated")

}
