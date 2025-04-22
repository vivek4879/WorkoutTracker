package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (app *application) AuthenticationHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://192.168.0.200:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&input)
	if err != nil {
		app.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}
	user, err := app.Models.UserModel.Query(input.Email)
	if err != nil {
		app.sendErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		log.Printf("User not found: %s", input.Email)
		return
	}

	match, err := app.PasswordHasher.Compare(input.Password, user.Password)
	if err != nil || !match {
		app.sendErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		log.Printf("Authentication failed for user: %s (match: %t, err: %v)", user.Email, match, err)
		return
	}

	log.Printf("User %s matches %t ", user.Email, match)
	//Create new random session token
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(48 * time.Hour)

	//// Delete existing session if exists
	if sess, err := app.Session(w, r); err == nil {
		if err := app.Models.UserModel.DeleteSession(sess); err != nil {
			log.Printf("Failed to delete old session: %v", err)
		}
	}
	err1 := app.Models.UserModel.InsertSession(user.ID, sessionToken, expiresAt)
	if err1 != nil {
		log.Println("Failed to insert session:", err1)
		app.sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
	response := map[string]string{
		"message":       "Authentication successful",
		"session_token": sessionToken,
		"user_id":       fmt.Sprintf("%d", user.ID),
	}
	log.Printf("Authentication successful for user: %s", user.Email)
	app.sendSuccessResponse(w, http.StatusOK, response)

}
