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

	dec := json.NewDecoder(r.Body)
	dec.Decode(&input)
	_, err1 := app.Models.UserModel.Query(input.Email)
	if err1 == nil {
		fmt.Printf("%s already exists\n", input.Email)
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	hashedPassword := hashing(input.Password)
	err := app.Models.UserModel.Insert(input.FirstName, input.LastName, input.Email, hashedPassword)
	fmt.Println("User created")

	if err != nil {
		fmt.Println("here")
		fmt.Println(err)
	}
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
