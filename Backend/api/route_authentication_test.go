package main

import (
	internal "WorkoutTracker/internal/database"
	"bytes"
	"encoding/json"
	"errors"
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

//Unhappy Path tests

func TestAuthenticationHandler_UserNotFound(t *testing.T) {
	mockUserModel := new(MockUserModel)
	mockHasher := new(MockPasswordHasher)

	testEmail := "test@email.com"
	testPassword := "securepassword"

	// Mock Query to simulate user not found
	mockUserModel.On("Query", testEmail).Return(nil, errors.New("user not found"))

	app := application{
		Models:         internal.Models{UserModel: mockUserModel},
		PasswordHasher: mockHasher,
	}

	reqBody := map[string]string{
		"email":    testEmail,
		"password": testPassword,
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.AuthenticationHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response JSON")
	}
	expected := "Invalid email or password"
	if response["error"] != expected {
		t.Errorf("Expected error %q, got %q", expected, response["error"])
	}

	mockUserModel.AssertExpectations(t)
}

func TestAuthenticationHandler_InvalidPassword(t *testing.T) {
	mockUserModel := new(MockUserModel)
	mockHasher := new(MockPasswordHasher)

	testEmail := "test@email.com"
	testPassword := "wrongpassword"
	correctPassword := "correctpassword"
	hashedPassword, _ := argon2id.CreateHash(correctPassword, argon2id.DefaultParams)

	mockUser := &internal.Users{
		ID:       1,
		Email:    testEmail,
		Password: hashedPassword,
	}

	mockUserModel.On("Query", testEmail).Return(mockUser, nil)
	mockHasher.On("Compare", testPassword, hashedPassword).Return(false, nil)

	app := application{
		Models:         internal.Models{UserModel: mockUserModel},
		PasswordHasher: mockHasher,
	}

	reqBody := map[string]string{
		"email":    testEmail,
		"password": testPassword,
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.AuthenticationHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response JSON")
	}
	expected := "Invalid email or password"
	if response["error"] != expected {
		t.Errorf("Expected error %q, got %q", expected, response["error"])
	}

	mockUserModel.AssertExpectations(t)
	mockHasher.AssertExpectations(t)
}
