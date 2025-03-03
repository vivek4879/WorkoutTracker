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

// TestAddWorkoutHandler tests adding a workout
func TestAddWorkoutHandler(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Mock session
	mockSession := &internal.Sessions{UserID: 1}

	// Mock workouts input
	workoutData := []internal.ExerciseData{
		{ExerciseId: 101, Sets: []internal.WorkoutSet{{SetNo: 1, Repetitions: 10, Weight: 50}}},
	}

	// Mock workout entry IDs returned from InsertWorkout
	mockWorkoutEntryIDs := []uint{201, 202}

	// Mock QuerySession to return a valid session
	mockUserModel.On("QuerySession", "mock-session-token").Return(mockSession, nil)

	// Mock InsertWorkout to return workout IDs
	mockUserModel.On("InsertWorkout", mockSession.UserID, workoutData).Return(mockWorkoutEntryIDs, nil)

	// Mock InsertWorkoutToUser to succeed
	mockUserModel.On("InsertWorkoutToUser", mockSession.UserID, mockWorkoutEntryIDs).Return(nil)

	// Setup Application with Mock Model
	mockModels := internal.Models{
		UserModel: mockUserModel,
	}
	app := application{Models: mockModels}

	// Create JSON request body
	reqBody := map[string]interface{}{
		"workouts": workoutData,
	}
	body, _ := json.Marshal(reqBody)

	// Create HTTP request
	req := httptest.NewRequest("POST", "/add-workout", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "mock-session-token"})

	// Create Response Recorder
	rec := httptest.NewRecorder()

	// Call the handler
	app.AddWorkoutHandler(rec, req)

	// Assertions
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rec.Code)
	}

	type Response struct {
		Message string `json:"message"`
	}

	var response Response
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v\nResponse Body: %s", err, rec.Body.String())
	}

	expectedMessage := "Workout added successfully"
	if response.Message != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, response.Message)
	}

	// Ensure mock expectations were met
	mockUserModel.AssertExpectations(t)
}
