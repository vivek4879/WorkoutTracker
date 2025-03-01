package main

import (
	"WorkoutTracker/internal/database"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHashing(t *testing.T) {
	password := "securepassword"
	hashed1 := Hashing(password)
	hashed2 := Hashing(password)
	if hashed1 == "" || hashed2 == "" {
		t.Error("Hashing function returned an empty string")
	}
	if hashed1 == password || hashed2 == password {
		t.Error("Hashing function returned plain text password")
	}
	if hashed1 == hashed2 {
		t.Error("Hashing function produced identical hashes for the same password (should be different due to salting)")
	}
}

func TestSession(t *testing.T) {
	mockUserModel := new(MockUserModel)
	validSession := database.Sessions{UserID: 1, Token: "valid-session"}
	mockUserModel.On("QuerySession", "valid-session").Return(&validSession, nil)

	mockApp := &application{
		Models: database.Models{
			UserModel: mockUserModel,
		},
	}
	req := httptest.NewRequest("GET", "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "valid-session"})
	rec := httptest.NewRecorder()
	session, err := mockApp.Session(rec, req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if session.Token != "valid-session" {
		t.Errorf("Expected token to be valid-session, got %s", session.Token)
	}
	reqNoCookie := httptest.NewRequest("GET", "/protected", nil)
	recNoCookie := httptest.NewRecorder()
	_, errNoCookie := mockApp.Session(recNoCookie, reqNoCookie)
	if errNoCookie == nil || errNoCookie.Error() != "No session token cookie" {
		t.Errorf("Expected no session token cookie, got %v", errNoCookie)
	}
	mockUserModel.On("QuerySession", "invalid-session").Return(nil, errors.New("Invalid session"))
	reqInvalid := httptest.NewRequest("GET", "/protected", nil)
	reqInvalid.AddCookie(&http.Cookie{Name: "session_token", Value: "invalid-session"})
	recInvalid := httptest.NewRecorder()
	_, errInvalid := mockApp.Session(recInvalid, reqInvalid)
	if errInvalid == nil || errInvalid.Error() != "Invalid or expired Session" {
		t.Errorf("Expected 'Invalid or expired Session' error, got %v", errInvalid)
	}
	mockUserModel.AssertExpectations(t)
}
