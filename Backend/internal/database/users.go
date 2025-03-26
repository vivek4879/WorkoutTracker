package database

import (
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func (u MyModel) GetAllExercises() ([]Exercises, error) {
	var exercises []Exercises
	result := u.db.Find(&exercises)
	return exercises, result.Error
}

func (u MyModel) UpsertUserBest(userID, exerciseID uint, weight, reps float64) error {
	best := UserBests{
		UserId:            userID,
		Ex_Id:             exerciseID,
		BestWeight:        weight,
		CorrespondingReps: reps,
	}

	//PostgreSQL upsert using GORM
	return u.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "userid"}, {Name: "ex_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"best_weight", "corresponding_reps"}),
	}).Create(&best).Error
}
func (u MyModel) QueryUserBest(UserId uint, Ex_Id uint) (bestweight float64, reps float64, err error) {
	var bestData UserBests
	res := u.db.Where("userid = ? AND ex_id = ?", UserId, Ex_Id).First(&bestData)
	if res.Error != nil {
		fmt.Println("Error fetching user's best:", res.Error)
		return 0., 0., res.Error
	}
	return bestData.BestWeight, bestData.CorrespondingReps, nil
}

func (u MyModel) InsertSession(Id uint, Token string, expiry time.Time) error {
	session := Sessions{
		UserID: Id,
		Token:  Token,
		Expiry: expiry,
	}
	// Use GORM's `OnConflict` to handle duplicate key violations
	res := u.db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "userid"}},                     // Conflict on `user_id`
			DoUpdates: clause.AssignmentColumns([]string{"token", "expiry"}), // Update these fields
		},
	).Create(&session)

	if res.Error != nil {
		fmt.Println("Error inserting/updating session:", res.Error)
		return res.Error
	}
	return nil
}

func (u MyModel) QueryLastWorkoutId(UserID uint) (uint, error) {
	var lastWorkoutId uint
	res := u.db.Model(&Workouts{}).
		Where("userid = ?", UserID).
		Select("COALESCE(MAX(workoutid), 0) AS workoutid").Scan(&lastWorkoutId)
	if res.Error != nil {
		return 0, res.Error
	}
	return lastWorkoutId, nil
}
func (u MyModel) InsertWorkout(UserID uint, workouts []ExerciseData) ([]uint, error) {

	lastWorkoutId, err := u.QueryLastWorkoutId(UserID)
	lastWorkoutId += 1
	if err != nil {
		return nil, err
	}

	var workoutBatch []Workouts

	for _, exercise := range workouts {
		var setNos, repetitions, weights []int64

		//convert sets to arrays
		for _, set := range exercise.Sets {
			setNos = append(setNos, int64(set.SetNo))
			repetitions = append(repetitions, int64(set.Repetitions))
			weights = append(weights, int64(set.Weight))
		}

		workout := Workouts{
			WorkoutId:         lastWorkoutId,
			UserId:            UserID,
			CurrentExerciseId: exercise.ExerciseId,
			SetNo:             pq.Int64Array(setNos),
			Repetitions:       pq.Int64Array(repetitions),
			Weights:           pq.Int64Array(weights),
			CreatedAt:         exercise.CreatedAt,
		}
		workoutBatch = append(workoutBatch, workout)
	}
	if err := u.db.Create(&workoutBatch).Error; err != nil {
		return nil, err
	}
	//retrieve inserted workout_entry_id's
	var insertWorkoutIDs []uint
	for _, workout := range workoutBatch {
		insertWorkoutIDs = append(insertWorkoutIDs, workout.WorkoutEntryID)
	}
	return insertWorkoutIDs, nil
}
func (u MyModel) UpdateMeasurements(userID uint, measurements Measurements) error {
	// Create a struct for the updated measurements
	updateFields := map[string]interface{}{}

	// Only update the fields that are non-zero (i.e., provided by the user)
	if measurements.Weight != nil {
		updateFields["weight"] = measurements.Weight
	}
	if measurements.Neck != nil {
		updateFields["neck"] = measurements.Neck
	}
	if measurements.Shoulders != nil {
		updateFields["shoulders"] = measurements.Shoulders
	}
	if measurements.Chest != nil {
		updateFields["chest"] = measurements.Chest
	}
	if measurements.LeftBicep != nil {
		updateFields["left_bicep"] = measurements.LeftBicep
	}
	if measurements.RightBicep != nil {
		updateFields["right_bicep"] = measurements.RightBicep
	}
	if measurements.UpperAbs != nil {
		updateFields["upper_abs"] = measurements.UpperAbs
	}
	if measurements.LowerAbs != nil {
		updateFields["lower_abs"] = measurements.LowerAbs
	}
	if measurements.Waist != nil {
		updateFields["waist"] = measurements.Waist
	}
	if measurements.Hips != nil {
		updateFields["hips"] = measurements.Hips
	}
	if measurements.LeftThigh != nil {
		updateFields["left_thigh"] = measurements.LeftThigh
	}
	if measurements.RightThigh != nil {
		updateFields["right_thigh"] = measurements.RightThigh
	}
	if measurements.LeftCalf != nil {
		updateFields["left_calf"] = measurements.LeftCalf
	}
	if measurements.RightCalf != nil {
		updateFields["right_calf"] = measurements.RightCalf
	}

	// Perform the update on the measurements table
	result := u.db.Model(&Measurements{}).Where("userid = ?", userID).Updates(updateFields)
	return result.Error
}

// GetMeasurements fetches the measurements for a given user
func (u MyModel) GetMeasurements(userID uint) (Measurements, error) {
	var measurements Measurements

	// Query the measurements table for the user's data
	result := u.db.Where("userid = ?", userID).First(&measurements)

	// If no record found, return empty Measurements struct
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return Measurements{}, nil // Return an empty struct if not found
		}
		return Measurements{}, result.Error
	}

	// Return the found measurements
	return measurements, nil
}

func (u MyModel) InsertWorkoutToUser(userID uint, workoutEntryIDs []uint) error {
	var workoutUserBatch []WorkoutToUser

	for _, workoutEntryID := range workoutEntryIDs {
		workoutUser := WorkoutToUser{
			UserID:                  userID,
			WorkoutEntryIDSecondary: workoutEntryID,
		}
		workoutUserBatch = append(workoutUserBatch, workoutUser)
	}
	if err := u.db.Create(&workoutUserBatch).Error; err != nil {
		return err
	}
	return nil
}
func (u MyModel) QuerySession(SessionToken string) (*Sessions, error) {
	var session Sessions
	res := u.db.Table("sessions").Where("Token = ?", SessionToken).First(&session)
	if res.Error != nil {
		return nil, res.Error
	}
	return &session, nil
}
func (u MyModel) Insert(FirstName string, LastName string, Email string, Password string) error {
	user := Users{
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email,
		Password:  Password,
	}
	res := u.db.Create(&user)
	if res.Error != nil {
		fmt.Println("Error while inserting into db")
		return res.Error
	}
	return nil
}

func (u MyModel) Query(Email string) (*Users, error) {
	var user Users
	res := u.db.Table("users").Where("email = ?", Email).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
func (u MyModel) QueryUserId(userID uint) (*Users, error) {
	var user Users
	res := u.db.Table("users").Where("userid = ?", userID).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u MyModel) DeleteSession(s Sessions) error {
	res := u.db.Delete(&s)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u MyModel) DeleteUser(s Users) error {
	res := u.db.Delete(&s)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
