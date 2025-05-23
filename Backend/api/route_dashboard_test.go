package main

import (
	internal "WorkoutTracker/internal/database"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
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
	const userID uint = 1

	// Create a new mock UserModel.
	mockUserModel := new(MockUserModel)

	// Prepare a mock session with UserID = 1.
	mockSession := &internal.Sessions{UserID: userID}
	exerciseID := uint(101)

	// Prepare workout input data.
	workoutData := []internal.ExerciseData{
		{
			ExerciseId: exerciseID,
			Sets: []internal.WorkoutSet{
				{SetNo: 1, Repetitions: 10, Weight: 50},
				{SetNo: 2, Repetitions: 8, Weight: 60}, // maximum weight set
			},
		},
	}

	// Prepare workout entry IDs returned from InsertWorkout.
	mockWorkoutEntryIDs := []uint{201, 202}

	// Set up expectations for each method call that will occur in the handler:
	mockUserModel.
		On("QuerySession", "mock-session-token").
		Return(mockSession, nil)
	mockUserModel.
		On("InsertWorkout", userID, workoutData).
		Return(mockWorkoutEntryIDs, nil)
	mockUserModel.
		On("InsertWorkoutToUser", userID, mockWorkoutEntryIDs).
		Return(nil)
	mockUserModel.
		On("QueryUserBest", userID, exerciseID).
		Return(40.0, 10.0, nil)
	mockUserModel.
		On("UpsertUserBest", userID, exerciseID, 60.0, 8.0).
		Return(nil)

	// For streak updates, assume that no streak exists yet.
	// The UpdateWorkoutStreak method is expected to:
	// 1. Call FetchStreakData(userID) → returns nil, nil.
	// 2. Then call UpsertStreak(newStreak) to create a new streak record.
	mockUserModel.
		On("FetchStreakData", userID).
		Return(nil, nil)
	mockUserModel.
		On("UpsertStreak", mock.AnythingOfType("*database.Streak")).
		Return(nil)
	// No expectation for UpdateWorkoutStreak because the implementation
	// calls UpsertStreak when no streak exists.

	// Set up the Application with the mocked model.
	mockModels := internal.Models{
		UserModel: mockUserModel,
	}
	app := application{Models: mockModels}

	// Create the JSON request body.
	reqBody := map[string]interface{}{
		"workouts": workoutData,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create an HTTP POST request with the JSON body and the session cookie.
	req := httptest.NewRequest("POST", "/add-workout", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "mock-session-token"})

	// Create a ResponseRecorder to capture the response.
	rec := httptest.NewRecorder()

	// Call the handler.
	app.AddWorkoutHandler(rec, req)

	// Validate the HTTP status code.
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rec.Code)
	}

	// Parse the response JSON.
	var respBody map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &respBody); err != nil {
		t.Fatalf("Failed to parse response JSON: %v\nResponse: %s", err, rec.Body.String())
	}

	// Check that the response message is as expected.
	expectedMessage := "Workout added successfully"
	if respBody["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, respBody["message"])
	}

	// Assert that FetchStreakData and UpsertStreak were called.
	mockUserModel.AssertCalled(t, "FetchStreakData", userID)
	mockUserModel.AssertCalled(t, "UpsertStreak", mock.AnythingOfType("*database.Streak"))

	// Verify that all expectations were met.
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

func TestGetStreakDataHandler(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Mock session data
	mockSession := &internal.Sessions{UserID: 3}
	// Mock streak data
	mockStreak := &internal.Streak{
		UserID:        3,
		CurrentStreak: 5,
		MaxStreak:     10,
	}

	// Mock QuerySession to return valid session
	mockUserModel.On("QuerySession", "mock-session-token").Return(mockSession, nil)
	// Mock FetchStreakData to return mock streak
	mockUserModel.On("FetchStreakData", mockSession.UserID).Return(mockStreak, nil)

	// Setup mock application with session handler
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

	// Create GET request
	req := httptest.NewRequest("GET", "/streak", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "mock-session-token"})
	rec := httptest.NewRecorder()

	// Call handler
	mockApp.GetStreakDataHandler(rec, req)

	// Assert status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Parse response
	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON response: %v\nBody: %s", err, rec.Body.String())
	}

	// Assert response content
	if response["user_id"] != float64(mockSession.UserID) {
		t.Errorf("Expected user_id %d, got %v", mockSession.UserID, response["user_id"])
	}
	if response["currentStreak"] != float64(mockStreak.CurrentStreak) {
		t.Errorf("Expected currentStreak %v, got %v", mockStreak.CurrentStreak, response["currentStreak"])
	}
	if response["maxStreak"] != float64(mockStreak.MaxStreak) {
		t.Errorf("Expected maxStreak %v, got %v", mockStreak.MaxStreak, response["maxStreak"])
	}

	mockUserModel.AssertExpectations(t)
}

// Unhappy Path tests
func TestGetAllExercisesHandler_DBError(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Simulate database failure
	mockUserModel.On("GetAllExercises").Return([]internal.Exercises(nil), errors.New("database error"))

	// Setup app
	mockModels := internal.Models{UserModel: mockUserModel}
	app := application{Models: mockModels}

	// Create request
	req := httptest.NewRequest("GET", "/exercises", nil)
	rec := httptest.NewRecorder()

	// Call the handler
	app.GetAllExercisesHandler(rec, req)

	// Check response code
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", rec.Code)
	}

	// Parse error response
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON error response: %v\nBody: %s", err, rec.Body.String())
	}

	expectedError := "Failed to retrieve exercises"
	if response["error"] != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, response["error"])
	}

	// Check mock expectations
	mockUserModel.AssertExpectations(t)
}

func TestDashboardHandler_InvalidTokenInDB(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Return error when QuerySession is called with token
	mockUserModel.On("QuerySession", "invalid-token").Return(nil, errors.New("invalid session"))

	mockApp := &MockApplication{
		application: application{
			Models: internal.Models{
				UserModel: mockUserModel,
			},
		},
		MockSession: func(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error) {
			return nil, errors.New("invalid session")
		},
	}

	req := httptest.NewRequest("GET", "/dashboard", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "invalid-token"})

	rec := httptest.NewRecorder()

	mockApp.DashboardHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized, got %d", rec.Code)
	}

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response")
	}

	if response["error"] != "Unauthorized: Invalid session" {
		t.Errorf("Expected error message 'Unauthorized: invalid session', got %q", response["error"])
	}
}

func TestDashboardHandler_MissingSessionCookie(t *testing.T) {
	mockUserModel := new(MockUserModel)

	mockApp := &MockApplication{
		application: application{
			Models: internal.Models{
				UserModel: mockUserModel,
			},
		},
		MockSession: func(w http.ResponseWriter, r *http.Request) (*internal.Sessions, error) {
			return nil, errors.New("no session token cookie")
		},
	}

	req := httptest.NewRequest("GET", "/dashboard", nil)
	rec := httptest.NewRecorder()

	mockApp.DashboardHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized, got %d", rec.Code)
	}

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response")
	}

	if response["error"] != "Unauthorized: Invalid session" {
		t.Errorf("Expected error 'Unauthorized: session token missing or invalid', got %q", response["error"])
	}
}

func TestAddWorkoutHandler_InsertWorkoutFails(t *testing.T) {
	const userID uint = 1
	mockUserModel := new(MockUserModel)
	mockSession := &internal.Sessions{UserID: userID}
	exerciseID := uint(101)

	workoutData := []internal.ExerciseData{
		{
			ExerciseId: exerciseID,
			Sets: []internal.WorkoutSet{
				{SetNo: 1, Repetitions: 10, Weight: 50},
			},
		},
	}

	// Set expectations
	mockUserModel.On("QuerySession", "mock-session-token").Return(mockSession, nil)
	mockUserModel.On("InsertWorkout", userID, workoutData).Return(nil, errors.New("insert workout failed"))

	app := application{
		Models: internal.Models{UserModel: mockUserModel},
	}

	reqBody := map[string]interface{}{
		"workouts": workoutData,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/add-workout", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "mock-session-token"})
	rec := httptest.NewRecorder()

	app.AddWorkoutHandler(rec, req)

	// Assert HTTP 500 status
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", rec.Code)
	}

	// Assert error response body
	var resp map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v\nBody: %s", err, rec.Body.String())
	}

	expected := "Internal Server Error"
	if resp["error"] != expected {
		t.Errorf("Expected error message %q, got %q", expected, resp["error"])
	}

	mockUserModel.AssertExpectations(t)
}

func TestAddHandler_QueryUserBestFails(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Mock session
	mockSession := &internal.Sessions{UserID: 2}
	mockExerciseID := uint(111)

	// Mock QuerySession to return valid session
	mockUserModel.On("QuerySession", "mock-session-token").Return(mockSession, nil)

	// Simulate DB error on QueryUserBest
	mockUserModel.On("QueryUserBest", mockSession.UserID, mockExerciseID).Return(0.0, 0.0, errors.New("db error"))

	// Setup mock application with session handler
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

	// Assert status code
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", rec.Code)
	}

	// Assert error message in response
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse error response: %v\nBody: %s", err, rec.Body.String())
	}

	expectedError := "Failed to retrieve user's best"
	if response["error"] != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, response["error"])
	}

	mockUserModel.AssertExpectations(t)
}

func TestGetStreakDataHandler_FetchStreakFails(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Mock session
	mockSession := &internal.Sessions{UserID: 3}

	// Mock QuerySession to return valid session
	mockUserModel.On("QuerySession", "mock-session-token").Return(mockSession, nil)

	// Simulate FetchStreakData failure
	mockUserModel.On("FetchStreakData", mockSession.UserID).Return(nil, errors.New("streak not found"))

	// Setup mock app with mocked session
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

	// Create GET request with session cookie
	req := httptest.NewRequest("GET", "/streak", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "mock-session-token"})
	rec := httptest.NewRecorder()

	// Call the handler
	mockApp.GetStreakDataHandler(rec, req)

	// Assert status code
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", rec.Code)
	}

	// Assert error response
	var response map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON response: %v\nBody: %s", err, rec.Body.String())
	}

	expectedError := "Internal Server Error"
	if response["error"] != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, response["error"])
	}

	mockUserModel.AssertExpectations(t)
}
