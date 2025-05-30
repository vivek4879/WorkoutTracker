package main

import (
	"WorkoutTracker/internal/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (app *application) UpdateMeasurementsHandler(w http.ResponseWriter, r *http.Request) {
	// Validate Session
	sess, err := app.Session(w, r)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		app.sendErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
		return
	}

	// Decode updated measurements from the request body
	var input database.Measurements
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Update measurements in the database
	err = app.Models.UserModel.UpdateMeasurements(sess.UserID, input)
	if err != nil {
		log.Printf("Error updating measurements: %v\n", err)
		app.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update measurements")
		return
	}

	// Send a success response
	app.sendSuccessResponse(w, http.StatusOK, map[string]string{"message": "Measurements updated successfully"})
}

func (app *application) GetMeasurementsHandler(w http.ResponseWriter, r *http.Request) {
	// Validate Session
	sess, err := app.Session(w, r)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		app.sendErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
		return
	}

	// Retrieve measurements for the user
	measurements, err := app.Models.UserModel.GetMeasurements(sess.UserID)

	if err != nil {
		log.Printf("Error fetching measurements: %v\n", err)
		app.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve measurements")
		return
	}
	// Send the measurements back in the response
	app.sendSuccessResponse(w, http.StatusOK, measurements)
}

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
	w.Header().Set("Access-Control-Allow-Origin", "http://192.168.0.200:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, session-token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
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

	//create slice to track best updates
	var bestUpdates []map[string]interface{}

	// Update bests for each exercise
	for _, workout := range input.Workouts {
		maxWeight := int64(0)
		correspondingReps := int64(0)

		//Find set with maximum weight
		for _, set := range workout.Sets {
			if set.Weight > maxWeight {
				maxWeight = set.Weight
				correspondingReps = set.Repetitions
			}
		}

		//Fetch current best
		currentBestWeight, _, err := app.Models.UserModel.QueryUserBest(sess.UserID, workout.ExerciseId)
		if err != nil {
			log.Printf("No existing best found or error: %v\n", err)
			//continue anyway to insert new best
		}

		//update if new best is greater
		if float64(maxWeight) > currentBestWeight {
			err = app.Models.UserModel.UpsertUserBest(sess.UserID, workout.ExerciseId, float64(maxWeight), float64(correspondingReps))
			if err != nil {
				log.Printf("Failed to update user best: %v\n", err)
				app.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update user's best performance")
				return
			}
			bestUpdates = append(bestUpdates, map[string]interface{}{
				"exercise_id":     workout.ExerciseId,
				"old_best_weight": currentBestWeight,
				"new_best_weight": float64(maxWeight),
				"new_best_reps":   correspondingReps,
			})
		}

	}

	// capture streak before updating
	oldStreak, _ := app.Models.UserModel.FetchStreakData(sess.UserID)

	err2 := app.UpdateWorkoutStreak(sess.UserID)
	if err2 != nil {
		app.sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// Capture new streak after update
	newStreak, _ := app.Models.UserModel.FetchStreakData(sess.UserID)
	response := map[string]interface{}{"message": "Workout added successfully",
		"updated_bests": bestUpdates,
		"Old_streak":    oldStreak,
		"new_streak":    newStreak,
	}
	app.sendSuccessResponse(w, http.StatusCreated, response)
}

// AddHandler Handler to return user's best in that particular exerciseId.
func (app *application) AddHandler(w http.ResponseWriter, r *http.Request) {
	// Validate Session
	sess, err := app.Session(w, r)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		app.sendErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
		return
	}

	// Parse exercise ID from query params
	exIDStr := r.URL.Query().Get("exercise_id")
	if exIDStr == "" {
		app.sendErrorResponse(w, http.StatusBadRequest, "Missing exercise_id parameter")
		return
	}
	exID, err := strconv.ParseUint(exIDStr, 10, 64)
	if err != nil {
		app.sendErrorResponse(w, http.StatusBadRequest, "Invalid exercise_id")
		return
	}

	// Query user's best
	bestWeight, reps, err := app.Models.UserModel.QueryUserBest(sess.UserID, uint(exID))
	if err != nil {
		log.Printf("Error querying user best: %v\n", err)
		app.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve user's best")
		return
	}

	// Send JSON response
	response := map[string]interface{}{
		"user_id":     sess.UserID,
		"exercise_id": exID,
		"best_weight": bestWeight,
		"reps":        reps,
	}
	app.sendSuccessResponse(w, http.StatusOK, response)
}
func (app *application) GetAllExercisesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://192.168.0.200:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	exercises, err := app.Models.UserModel.GetAllExercises()
	if err != nil {
		log.Printf("Error fetching exercises: %v\n", err)
		app.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve exercises")
		return
	}

	// Include ID, Name, and URL in the response
	type ExerciseDTO struct {
		ID   uint   `json:"exercise_id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	var response []ExerciseDTO
	for _, ex := range exercises {
		response = append(response, ExerciseDTO{
			ID:   ex.ExerciseId,
			Name: ex.ExerciseName,
			URL:  ex.ExerciseImageURL,
		})
	}

	app.sendSuccessResponse(w, http.StatusOK, response)
}
func (app *application) GetStreakDataHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := app.Session(w, r)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		app.sendErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Invalid session")
		return
	}
	StreakData, err := app.Models.UserModel.FetchStreakData(sess.UserID)
	if err != nil {
		app.sendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	response := map[string]interface{}{
		"user_id":       sess.UserID,
		"currentStreak": StreakData.CurrentStreak,
		"maxStreak":     StreakData.MaxStreak,
	}
	app.sendSuccessResponse(w, http.StatusOK, response)

}
