package main

import (
	"errors"
	"net/http"
	"time"
)

func (app *application) CheckCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No session token cookie", http.StatusUnauthorized)
			return "", errors.New("No session token cookie")
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return "", err
	}
	return cookie.Value, nil
}

func (app *application) CheckToken(w http.ResponseWriter, sessionToken string) (*string, error) {
	s, err := app.Models.UserModel.QUERYSESSION(sessionToken)
	if err != nil {
		http.Error(w, "Unauthorized: Session not found", http.StatusUnauthorized)
		return nil, err
	}
	if s.Expiry.Before(time.Now()) {
		err1 := app.Models.UserModel.DeleteSession(*s)
		if err1 != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil, err1
		}
		http.Error(w, "Unauthorized: Session expired", http.StatusUnauthorized)
		return nil, errors.New("session expired")
	}
	return &s.Token, nil
}
