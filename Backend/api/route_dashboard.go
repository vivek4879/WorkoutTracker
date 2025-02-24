package main

import (
	"log"
	"net/http"
)

func (app *application) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	//Validate Session
	sess, err := app.Session(w, r)
	if err != nil {
		log.Printf("Error getting session: %v", err)
		app.sendErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
		return
	}
	//Log Successful access
	log.Printf("User %s accessed the Dashboard", sess.UserID)

	// Return JSON response
	response := map[string]string{
		"message": "Welcome to the dashboard page",
	}
	log.Printf("Response from Dashboard: %s", response)
	app.sendSuccessResponse(w, http.StatusOK, response)

}

func (app *application) AddExerciseHandler(w http.ResponseWriter, r *http.Request) {
	//Validate Session
	sess, err := app.Session(w, r)
	if err != nil {
		log.Printf("Error getting session: %v", err)
		app.sendErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
		return
	}
	log.Printf("User %s added exercise", sess.UserID)

}
