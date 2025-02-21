package main

import (
	"encoding/json"
	"fmt"
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
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No session token cookie", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	s, err2 := app.Models.UserModel.QUERYSESSION(sessionToken)
	if err2 != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	err3 := app.Models.UserModel.DeleteSession(*s)
	if err3 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	//http.Error(w, "Unauthorized:Session expired", http.StatusUnauthorized)

	fmt.Println("User successfully logged out")

	return
}

func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := app.CheckCookie(w, r)
	if err != nil {
		http.Error(w, "Unauthorized: Session not found", http.StatusUnauthorized)
		return
	}
	session, err2 := app.Models.UserModel.QUERYSESSION(sessionToken)
	if err2 != nil {
		fmt.Println(err2)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	user, err3 := app.Models.UserModel.QueryUserId(session.UserID)
	if err3 != nil {
		fmt.Println(err3)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err4 := app.Models.UserModel.DeleteUser(*user)
	if err4 != nil {
		fmt.Println(err4)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err5 := app.Models.UserModel.DeleteSession(*session)
	if err5 != nil {
		fmt.Println(err5)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("User successfully deleted")
}
