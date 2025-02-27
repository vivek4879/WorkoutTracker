package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	internal "WorkoutTracker/internal/database"
)

// Mock application with a session method
//type MockApplication struct {
//	application
//	MockSession func(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error)
//}
//
//func (m *MockApplication) Session(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error) {
//	return m.MockSession(w, r)
//}

func TestDashboardHandler(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Mock a valid session
	validSession := &internal.Sessions{UserID: 1, Token: "mock-session-token"}
	//  Mock `QuerySession` to return a valid session
	mockUserModel.On("QuerySession", "mock-session-token").Return(validSession, nil)
	//  Mock Application with valid session handling
	mockApp := &MockApplication{
		application: application{
			Models: internal.Models{
				UserModel: mockUserModel,
			},
		},
		MockSession: func(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error) {
			return validSession, nil // Return valid session
		},
	}

	//  Create HTTP request
	req := httptest.NewRequest("GET", "/dashboard", bytes.NewReader([]byte{}))
	req.AddCookie(&http.Cookie{Name: "session_token", Value: validSession.Token})

	//  Create Response Recorder
	rec := httptest.NewRecorder()

	//  Call the Dashboard handler
	mockApp.DashboardHandler(rec, req)

	//  Assertions: Expected 200 OK response
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	//  Verify JSON Response
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response JSON")
	}

	expectedMessage := "Welcome to the dashboard page"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, response["message"])
	}

	//  Case 2: Mock an invalid session
	mockApp.MockSession = func(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error) {
		return nil, http.ErrNoCookie // Simulate invalid session
	}

	//  Create request for unauthorized access
	reqUnauthorized := httptest.NewRequest("GET", "/dashboard", bytes.NewReader([]byte{}))

	//  Create Response Recorder
	recUnauthorized := httptest.NewRecorder()

	//  Call the Dashboard handler with unauthorized request
	mockApp.DashboardHandler(recUnauthorized, reqUnauthorized)

	//  Assertions: Expected 401 Unauthorized response
	if recUnauthorized.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", recUnauthorized.Code)
	}

	log.Println(" Dashboard handler test passed successfully.")
}
