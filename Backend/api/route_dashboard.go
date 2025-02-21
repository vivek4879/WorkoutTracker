package main

import (
	"fmt"
	"net/http"
)

func (app *application) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := app.CheckCookie(w, r)
	if err != nil {
		return // Exit early if error occurs
	}

	_, err3 := app.CheckToken(w, sessionToken)
	if err3 != nil {
		return // Exit early if token is invalid
	}

	fmt.Println("Welcome to the home page")
	_, writeErr := w.Write([]byte("Welcome to the home page!\n"))
	if writeErr != nil {
		fmt.Println("Error writing response:", writeErr)
	}
}
