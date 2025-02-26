package main

import (
	internal "WorkoutTracker/internal/database"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	mockUserModel := new(MockUserModel) // create a mock user model

	//Expect insert to be called with these parameters and return nil if successful
	mockUserModel.On("Insert", "Vivek", "Aher", "Vivek@gmail.com", "securepassword").Return(nil)

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
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
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
