package main

import (
	internal "WorkoutTracker/internal/database"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	mockUserModel := new(MockUserModel) // create a mock user model

	// Expect Query to be called before Insert
	mockUserModel.On("Query", "test@email.com").Return(nil, errors.New("user not found")) // Simulate "User Not Found"
	//Expect insert to be called with these parameters and return nil if successful
	mockUserModel.On("Insert", "Vivek", "Aher", "test@email.com", mock.Anything).Return(nil)
	// Mocking other methods that might be used
	//mockUserModel.On("InsertSession", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	//mockUserModel.On("QuerySession", mock.Anything).Return(nil, nil)
	//mockUserModel.On("DeleteSession", mock.Anything).Return(nil)
	//mockUserModel.On("QueryUserId", mock.Anything).Return(nil, nil)
	//mockUserModel.On("DeleteUser", mock.Anything).Return(nil)
	//mockUserModel.On("InsertWorkout", mock.Anything, mock.Anything).Return([]uint{1, 2}, nil)
	//mockUserModel.On("InsertWorkoutToUser", mock.Anything, mock.Anything).Return(nil)

	// Create an instance of internal.Models, replacing userModel with mock
	mockModels := internal.Models{
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
