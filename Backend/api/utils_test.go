package main

import (
	"WorkoutTracker/internal/database"
	"encoding/json"
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
	if errNoCookie == nil || errNoCookie.Error() != "no session token cookie" {
		t.Errorf("Expected no session token cookie, got %v", errNoCookie)
	}
	mockUserModel.On("QuerySession", "invalid-session").Return(nil, errors.New("invalid session"))
	reqInvalid := httptest.NewRequest("GET", "/protected", nil)
	reqInvalid.AddCookie(&http.Cookie{Name: "session_token", Value: "invalid-session"})
	recInvalid := httptest.NewRecorder()
	_, errInvalid := mockApp.Session(recInvalid, reqInvalid)
	if errInvalid == nil || errInvalid.Error() != "invalid or expired Session" {
		t.Errorf("Expected 'invalid or expired Session' error, got %v", errInvalid)
	}
	mockUserModel.AssertExpectations(t)
}

func TestSendErrorResponse(t *testing.T) {
	mockApp := &application{}

	rec := httptest.NewRecorder()
	mockApp.sendErrorResponse(rec, http.StatusForbidden, "Access Denied")

	if rec.Code != http.StatusForbidden {
		t.Errorf("Expected status 403, got %d", rec.Code)
	}

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse JSON response")
	}

	expectedError := "Access Denied"
	if response["error"] != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, response["error"])
	}
}

func TestSendSuccessResponse(t *testing.T) {
	mockApp := &application{}

	rec := httptest.NewRecorder()
	responseData := map[string]string{"message": "Operation successful"}

	mockApp.sendSuccessResponse(rec, http.StatusOK, responseData)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse JSON response")
	}

	expectedMessage := "Operation successful"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, response["message"])
	}
}
