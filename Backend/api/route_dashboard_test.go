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
//
//	type MockApplication struct {
//		application
//		MockSession func(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error)
//	}
//
//	func (m *MockApplication) Session(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error) {
//		return m.MockSession(w, r)
//	}
func TestGetAllExercisesHandler(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Sample mock data
	exercises := []internal.Exercises{
		{ExerciseId: 101, ExerciseName: "Bench Press", ExerciseImageURL: "https://s3.bucket/bench-press.png"},
		{ExerciseId: 102, ExerciseName: "Deadlift", ExerciseImageURL: "https://s3.bucket/deadlift.png"},
	}

	// Expect the GetAllExercises call and return mock data
	mockUserModel.On("GetAllExercises").Return(exercises, nil)

	// Setup app
	mockModels := internal.Models{UserModel: mockUserModel}
	app := application{Models: mockModels}

	// Create request
	req := httptest.NewRequest("GET", "/exercises", nil)
	rec := httptest.NewRecorder()

	// Call the handler
	app.GetAllExercisesHandler(rec, req)

	// Assertions
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response []map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v\nBody: %s", err, rec.Body.String())
	}

	// Basic checks
	if len(response) != 2 {
		t.Errorf("Expected 2 exercises, got %d", len(response))
	}

	if response[0]["name"] != "Bench Press" || response[0]["url"] != "https://s3.bucket/bench-press.png" {
		t.Errorf("Unexpected first exercise: %+v", response[0])
	}

	// Check mock expectations
	mockUserModel.AssertExpectations(t)
}

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
	exerciseID := uint(101)

	// Mock workouts input
	workoutData := []internal.ExerciseData{
		{
			ExerciseId: exerciseID,
			Sets: []internal.WorkoutSet{
				{SetNo: 1, Repetitions: 10, Weight: 50},
				{SetNo: 2, Repetitions: 8, Weight: 60}, // max weight
			},
		},
	}

	// Mock workout entry IDs returned from InsertWorkout
	mockWorkoutEntryIDs := []uint{201, 202}

	// Mock QuerySession to return a valid session
	mockUserModel.On("QuerySession", "mock-session-token").Return(mockSession, nil)

	// Mock InsertWorkout to return workout IDs
	mockUserModel.On("InsertWorkout", mockSession.UserID, workoutData).Return(mockWorkoutEntryIDs, nil)

	// Mock InsertWorkoutToUser to succeed
	mockUserModel.On("InsertWorkoutToUser", mockSession.UserID, mockWorkoutEntryIDs).Return(nil)

	// Mock QueryUserBest to return a lower best so the new one is inserted
	mockUserModel.On("QueryUserBest", mockSession.UserID, exerciseID).Return(40.0, 10.0, nil)

	// Mock UpsertUserBest to succeed
	mockUserModel.On("UpsertUserBest", mockSession.UserID, exerciseID, 60.0, 8.0).Return(nil)

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

	// Ensure all mock expectations were met
	mockUserModel.AssertExpectations(t)
}

func TestAddHandler(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Mock session
	mockSession := &internal.Sessions{UserID: 2}
	mockExerciseID := uint(111)
	mockBestWeight := 100.0
	mockReps := 8.0

	// Mock QuerySession to return valid session
	mockUserModel.On("QuerySession", "mock-session-token").Return(mockSession, nil)

	// Mock QueryUserBest to return the expected best weight and reps
	mockUserModel.On("QueryUserBest", mockSession.UserID, mockExerciseID).Return(mockBestWeight, mockReps, nil)

	// Setup mock application with mock session handler
	mockApp := &MockApplication{
		application: application{
			Models: internal.Models{
				UserModel: mockUserModel,
			},
		},
		MockSession: func(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error) {
			return mockSession, nil
		},
	}

	// Create request with query param: ?exercise_id=111
	req := httptest.NewRequest("GET", "/user-best?exercise_id=111", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "mock-session-token"})

	// Create response recorder
	rec := httptest.NewRecorder()

	// Call the handler
	mockApp.AddHandler(rec, req)

	// Assert response status
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Assert response JSON
	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v\nBody: %s", err, rec.Body.String())
	}

	if response["user_id"] != float64(mockSession.UserID) {
		t.Errorf("Expected user_id %v, got %v", mockSession.UserID, response["user_id"])
	}
	if response["exercise_id"] != float64(mockExerciseID) {
		t.Errorf("Expected exercise_id %v, got %v", mockExerciseID, response["exercise_id"])
	}
	if response["best_weight"] != mockBestWeight {
		t.Errorf("Expected best_weight %v, got %v", mockBestWeight, response["best_weight"])
	}
	if response["reps"] != mockReps {
		t.Errorf("Expected reps %v, got %v", mockReps, response["reps"])
	}

	// Ensure all mock expectations were met
	mockUserModel.AssertExpectations(t)
}
