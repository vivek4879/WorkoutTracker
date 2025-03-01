package main

import (
	"WorkoutTracker/internal/database"
	"encoding/json"
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
	log.Printf("User %U accessed the Dashboard\n", sess.UserID)

	// Return JSON response
	response := map[string]string{
		"message": "Welcome to the dashboard page",
	}
	log.Printf("Response from Dashboard: %s\n", response)
	app.sendSuccessResponse(w, http.StatusOK, response)

}

func (app *application) AddWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	//Validate Session
	sess, err := app.Session(w, r)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		app.sendErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
		return
	}
	log.Printf("User %U added exercise\n", sess.UserID)

	//decode workout data from request body
	var input struct {
		Workouts []database.ExerciseData `json:"workouts"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	// insert workout with validated userid and get workout_entry_ids

	workoutEntryIDs, err := app.Models.UserModel.InsertWorkout(sess.UserID, input.Workouts)
	if err != nil {
		app.sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	//insert into workoutToUser Table
	err1 := app.Models.UserModel.InsertWorkoutToUser(sess.UserID, workoutEntryIDs)
	if err1 != nil {
		app.sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response := map[string]string{"message": "Workout added successfully"}
	app.sendSuccessResponse(w, http.StatusCreated, response)
}
