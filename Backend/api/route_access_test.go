package main

import (
	"WorkoutTracker/internal/database"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	mockUserModel := new(MockUserModel) // create a mock user model

	// Expect Query to be called before Insert
	mockUserModel.On("Query", "test@email.com").Return(nil, errors.New("user not found")) // Simulate "User Not Found"
	//Expect insert to be called with these parameters and return nil if successful
	mockUserModel.On("Insert", "Vivek", "Aher", "test@email.com", mock.Anything).Return(nil)

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

	//converting request body to JSON
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Error("Failed to marshal request body")
	}
	req := httptest.NewRequest("POST", "/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.signupHandler(rec, req)

	//check response
	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response JSON")
	}

	expectedMessage := "User successfully created"
	if response["message"] != expectedMessage {
		t.Errorf("Expected response message %q, got %q", expectedMessage, response["message"])
	}

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

//func TestSignupHandler(t *testing.T) {
//	// setting up test application
//	app := application{Models: NewModels(setupTestDB())}
//
//	//mock user signup request
//	reqBody := map[string]string{
//		"firstname": "Vivek",
//		"lastname":  "Aher",
//		"email":     "test@email.com",
//		"password":  "securepassword",
//	}
//
//	//converting request body to JSON
//	body, _ := json.Marshal(reqBody)
//	//create mock HTTP request using Go's httptest package
//	req := httptest.NewRequest("POST", "/signup", bytes.NewReader(body))
//	req.Header.Set("Content-Type", "application/json")
//	rec := httptest.NewRecorder()
//
//	app.signupHandler(rec, req)
//	if rec.Code != http.StatusOK {
//		t.Errorf("Expected status 200, got %d", rec.Code)
//	}
//}
