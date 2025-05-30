package main

import (
	"WorkoutTracker/internal/database"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	mockUserModel := new(MockUserModel) // create a mock user model

	// Expect Query to be called before Insert
	mockUserModel.On("Query", "test@email.com").Return(nil, errors.New("user not found")) // Simulate "User Not Found"

	// Expect insert to be called with these parameters and return nil if successful
	mockUserModel.On("Insert", "Vivek", "Aher", "test@email.com", mock.Anything).Return(nil)

	// Expect InsertBlankMeasurements to be called to create blank measurements for the user
	mockUserModel.On("InsertBlankMeasurements", mock.Anything).Return(nil)

	mockUserModel.On("GetUserIDByEmail", "test@email.com").Return(uint(1), nil)

	// Create an instance of internal.Models, replacing userModel with mock
	mockModels := database.Models{
		UserModel: mockUserModel, // inject mock user model
	}
	app := application{Models: mockModels} // Inject mock model into app

	reqBody := map[string]string{
		"firstname": "Vivek",
		"lastname":  "Aher",
		"email":     "test@email.com",
		"password":  "securepassword",
	}

	// Converting request body to JSON
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Error("Failed to marshal request body")
	}
	req := httptest.NewRequest("POST", "/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Call signupHandler
	app.signupHandler(rec, req)

	// Check response
	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response JSON")
	}

	expectedMessage := "User successfully created"
	if response["message"] != expectedMessage {
		t.Errorf("Expected response message %q, got %q", expectedMessage, response["message"])
	}

	// Ensure InsertBlankMeasurements was called (create blank measurements entry)
	mockUserModel.AssertExpectations(t)
}

// Mock session struct
// Instead of modifying application, we create a mock version inside our test.
// This overrides Session() only for the test, without causing conflicts
// It uses a function MockSession that we can set dynamically in the test
type MockApplication struct {
	application
	MockSession func(w http.ResponseWriter, r *http.Request) (*database.Sessions, error)
}

func (m *MockApplication) Session(w http.ResponseWriter, r *http.Request) (*database.Sessions, error) {
	return m.MockSession(w, r)
}

func TestDeleteHandler(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Create a mock session
	mockSession := &database.Sessions{
		UserID: 1,
		Token:  "mock-session-token",
	}

	// Create a mock user
	mockUser := &database.Users{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Password:  "hashedpassword",
	}

	// Mock session retrieval
	mockUserModel.On("QuerySession", "mock-session-token").Return(mockSession, nil)

	// Mock user retrieval
	mockUserModel.On("QueryUserId", mockSession.UserID).Return(mockUser, nil)

	// Mock user deletion
	mockUserModel.On("DeleteUser", *mockUser).Return(nil)

	// Mock session deletion
	mockUserModel.On("DeleteSession", *mockSession).Return(nil)

	// Setup application with mock model
	mockModels := database.Models{
		UserModel: mockUserModel,
	}
	app := application{Models: mockModels}

	// Create HTTP request with session token
	req := httptest.NewRequest("DELETE", "/delete", bytes.NewReader([]byte{}))
	req.AddCookie(&http.Cookie{Name: "session_token", Value: mockSession.Token})

	// Create Response Recorder
	rec := httptest.NewRecorder()

	// Call the deleteHandler function
	app.deleteHandler(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code, "Expected status 200 OK")

	// Verify JSON Response
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to parse response JSON")
	assert.Equal(t, "User successfully deleted", response["message"])

	// Ensure mock expectations were met
	mockUserModel.AssertExpectations(t)
}

//Unhappy path tests

func TestSignupHandler_UserAlreadyExists(t *testing.T) {
	mockUserModel := new(MockUserModel)

	// Simulate user already exists by returning some user data instead of error
	mockUserModel.On("Query", "test@email.com").Return(&database.Users{
		ID:        1,
		FirstName: "Existing",
		LastName:  "User",
		Email:     "test@email.com",
	}, nil)

	mockModels := database.Models{
		UserModel: mockUserModel,
	}
	app := application{Models: mockModels}

	reqBody := map[string]string{
		"firstname": "Vivek",
		"lastname":  "Aher",
		"email":     "test@email.com",
		"password":  "securepassword",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Error("Failed to marshal request body")
	}
	req := httptest.NewRequest("POST", "/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Call signupHandler
	app.signupHandler(rec, req)

	// Check the HTTP status code
	if rec.Code != http.StatusConflict {
		t.Errorf("Expected status code %d, got %d", http.StatusConflict, rec.Code)
	}

	// Check response body message
	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response JSON")
	}

	expectedMessage := "Email already exists"
	if response["error"] != expectedMessage {
		t.Errorf("Expected error message %q, got %q", expectedMessage, response["error"])
	}

	// Assert expectations
	mockUserModel.AssertExpectations(t)
}

func TestDeleteHandler_SessionNotFound(t *testing.T) {
	mockUserModel := new(MockUserModel)
	mockUserModel.On("QuerySession", "mock-session-token").Return(nil, errors.New("session not found"))

	app := application{Models: database.Models{UserModel: mockUserModel}}

	req := httptest.NewRequest("DELETE", "/delete", nil)
	req.AddCookie(&http.Cookie{Name: "session_token", Value: "mock-session-token"})
	rec := httptest.NewRecorder()

	app.deleteHandler(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Unauthorized: Invalid session", response["error"])

	mockUserModel.AssertExpectations(t)
}

func TestDeleteHandler_MissingSessionToken(t *testing.T) {
	mockUserModel := new(MockUserModel)
	app := application{Models: database.Models{UserModel: mockUserModel}}

	req := httptest.NewRequest("DELETE", "/delete", nil)
	rec := httptest.NewRecorder()

	app.deleteHandler(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Unauthorized: Invalid session", response["error"])
}
