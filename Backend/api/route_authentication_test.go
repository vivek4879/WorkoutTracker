package main

import (
	internal "WorkoutTracker/internal/database"
	"bytes"
	"encoding/json"
	"github.com/alexedwards/argon2id"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockPasswordHasher struct {
	mock.Mock
}

func (m *MockPasswordHasher) Compare(password, hash string) (bool, error) {
	args := m.Called(password, hash)
	return args.Bool(0), args.Error(1)
}
func TestAuthenticationHandler(t *testing.T) {
	mockUserModel := new(MockUserModel)
	mockHasher := new(MockPasswordHasher)

	// Test User Data
	testEmail := "test@email.com"
	testPassword := "securepassword"
	hashedPassword, _ := argon2id.CreateHash(testPassword, argon2id.DefaultParams) // Mock hashing
	mockUser := &internal.Users{
		ID:       1,
		Email:    testEmail,
		Password: hashedPassword,
	}

	// Mock Query: User Exists
	mockUserModel.On("Query", testEmail).Return(mockUser, nil)

	// Mock Password Check
	mockHasher.On("Compare", testPassword, hashedPassword).Return(true, nil)

	// Mock UUID Generation
	sessionToken := "mock-session-token"

	// Mock InsertSession
	var generatedToken string
	mockUserModel.On("InsertSession", mockUser.ID, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			generatedToken = args.String(1) // Capture the session token
		}).
		Return(nil)

	// Setup Application with Mock Model
	mockModels := internal.Models{
		UserModel: mockUserModel,
	}
	app := application{Models: mockModels,
		PasswordHasher: mockHasher}

	// Request Body
	reqBody := map[string]string{
		"email":    testEmail,
		"password": testPassword,
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Call the handler
	app.AuthenticationHandler(rec, req)

	// Assertions
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Verify JSON Response
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response JSON")
	}

	expectedToken := sessionToken
	if response["session_token"] != generatedToken {
		t.Errorf("Expected session token %q, got %q", expectedToken, response["session_token"])
	}

	expectedMessage := "Authentication successful"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, response["message"])
	}

	// Ensure mock expectations were met
	mockUserModel.AssertExpectations(t)
	mockHasher.AssertExpectations(t)
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
