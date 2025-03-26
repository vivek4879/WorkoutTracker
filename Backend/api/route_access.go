package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (app *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	// Decode the request body
	dec := json.NewDecoder(r.Body)
	dec.Decode(&input)

	// Check if the email already exists using the existing Query method
	_, err1 := app.Models.UserModel.Query(input.Email)
	if err1 == nil {
		// If email exists, return a conflict error
		fmt.Printf("%s already exists\n", input.Email)
		app.sendErrorResponse(w, http.StatusConflict, "Email already exists")
		return
	}

	// Hash the user's password
	hashedPassword := Hashing(input.Password)

	// Insert the user into the database
	err := app.Models.UserModel.Insert(input.FirstName, input.LastName, input.Email, hashedPassword)
	if err != nil {
		// If user insert fails, return an internal server error
		fmt.Printf("User %s not created\n", input.Email)
		app.sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Get the newly created user's ID using the email
	userID, err := app.Models.UserModel.GetUserIDByEmail(input.Email)
	if err != nil {
		// If there's an error fetching user ID, return an internal server error
		fmt.Printf("Error fetching userID for %s\n", input.Email)
		app.sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Insert a blank entry into the measurements table for the new user
	err = app.Models.UserModel.InsertBlankMeasurements(userID)
	if err != nil {
		// If there was an error inserting into measurements, return an internal server error
		fmt.Printf("Error inserting blank measurements for user %d\n", userID)
		app.sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Send success response
	response := map[string]string{
		"message": "User successfully created",
	}
	app.sendSuccessResponse(w, http.StatusOK, response)
}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve session token from cookie
	sess, err := app.Session(w, r)
	if err != nil {
		app.sendErrorResponse(w, http.StatusUnauthorized, err.Error())
		log.Printf("logout failed:%v", err)
		return
	}
	//Delete session
	err3 := app.Models.UserModel.DeleteSession(sess)
	if err3 != nil {
		app.sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete session")
		log.Printf("logout failed:Error deleting session%v", err3)
		return
	}

	log.Printf("logout success")
	app.sendSuccessResponse(w, http.StatusOK, map[string]string{"message": "User successfully logged out"})
}

func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := app.Session(w, r)
	if err != nil {
		app.sendErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
		log.Printf("Error getting session: %v", err)
		return
	}

	user, err3 := app.Models.UserModel.QueryUserId(sess.UserID)
	if err3 != nil {
		app.sendErrorResponse(w, http.StatusUnauthorized, "Failed to retrieve user")
		log.Printf("Error getting user: %v", err3)

		return
	}
	err4 := app.Models.UserModel.DeleteUser(*user)
	if err4 != nil {
		app.sendErrorResponse(w, http.StatusUnauthorized, "Failed to delete user")
		log.Printf("Error deleting user: %v", err4)
		return
	}
	err5 := app.Models.UserModel.DeleteSession(sess)
	if err5 != nil {
		app.sendErrorResponse(w, http.StatusUnauthorized, "Failed to delete session")
		log.Printf("Error deleting session: %v", err5)
		return
	}
	log.Println("User successfully deleted")
	//create a json response object, we will convert this map to json before sending to client
	response := map[string]string{
		"message": "User successfully deleted",
	}
	app.sendSuccessResponse(w, http.StatusOK, response)

}
