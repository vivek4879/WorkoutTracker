package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (app *application) AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&input)
	if err != nil {
		http.Error(w, `{"error" : "Invalid JSON input"}`, http.StatusBadRequest)
		return
	}
	user, err := app.Models.UserModel.Query(input.Email)
	if err != nil {
		fmt.Println(err)
		log.Printf("User not found: %s", input.Email)
		http.Error(w, `{"error" : "Invalid email or password"}`, http.StatusUnauthorized)
		return
	}
	match, err := argon2id.ComparePasswordAndHash(input.Password, user.Password)
	if err != nil {
		log.Printf("Incorrect Password for user: %s ", user.Email)
		http.Error(w, `{"error": "Invalid email or password"}`, http.StatusUnauthorized)
		return
	}
	log.Printf("User %s matches %t ", user.Email, match)
	//Create new random session token
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(48 * time.Hour)

	err1 := app.Models.UserModel.InsertSession(user.ID, sessionToken, expiresAt)
	if err1 != nil {
		log.Println("Failed to insert session:", err1)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
	}
	cookie := http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	}
	http.SetCookie(w, &cookie)
	response := map[string]string{
		"message":       "Authentication successful",
		"session_token": sessionToken,
		"user_id":       fmt.Sprintf("%d", user.ID),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
