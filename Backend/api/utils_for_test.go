package main

import (
	"WorkoutTracker/internal/database"
	"github.com/stretchr/testify/mock"
	"time"
)

// simulate database
type MockUserModel struct {
	mock.Mock
}

func (m *MockUserModel) QueryUserBest(UserId uint, Ex_Id uint) (float64, float64, error) {
	args := m.Called(UserId, Ex_Id)
	return args.Get(0).(float64), args.Get(1).(float64), args.Error(2)
}

// Mock function for user signup
func (m *MockUserModel) Insert(firstname, lastname, email, password string) error {
	args := m.Called(firstname, lastname, email, password)
	return args.Error(0)
}

// Mock Query function for checking if user exists
func (m *MockUserModel) Query(email string) (*database.Users, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*database.Users), args.Error(1)
}

// Mock InsertSession function (for session management)
func (m *MockUserModel) InsertSession(Id uint, Token string, expiry time.Time) error {
	args := m.Called(Id, Token, expiry)
	return args.Error(0)
}

// Mock QuerySession function (for retrieving a session)
func (m *MockUserModel) QuerySession(SessionToken string) (*database.Sessions, error) {
	args := m.Called(SessionToken)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*database.Sessions), args.Error(1)
}

// Mock DeleteSession function
func (m *MockUserModel) DeleteSession(s database.Sessions) error {
	args := m.Called(s)
	return args.Error(0)
}

// Mock QueryUserId function
func (m *MockUserModel) QueryUserId(userID uint) (*database.Users, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*database.Users), args.Error(1)
}

// Mock DeleteUser function
func (m *MockUserModel) DeleteUser(u database.Users) error {
	args := m.Called(u)
	return args.Error(0)
}

// Mock InsertWorkout function
func (m *MockUserModel) InsertWorkout(UserID uint, workouts []database.ExerciseData) ([]uint, error) {
	args := m.Called(UserID, workouts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]uint), args.Error(1)
}

// Mock InsertWorkoutToUser function
func (m *MockUserModel) InsertWorkoutToUser(userID uint, workoutEntryIDs []uint) error {
	args := m.Called(userID, workoutEntryIDs)
	return args.Error(0)
}

func (m *MockUserModel) UpsertUserBest(userID uint, Ex_Id uint, weight float64, reps float64) error {
	args := m.Called(userID, Ex_Id, weight, reps)
	return args.Error(0)
}

//func setupTestDB() *gorm.DB {
//	dsn := "host=localhost user=test password=test dbname=testdb port=5433 sslmode=disable"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %v", err)
//	}
//	err1 := db.AutoMigrate(&database.Users{}, &database.Sessions{}, &database.Exercises{}, &database.WorkoutToUser{}, &database.Workouts{})
//	if err1 != nil {
//		log.Fatalf("Failed to migrate database: %v", err1)
//	}
//	//clean all tables before each test
//	db.Exec("TRUNCATE users,sessions,exercises,workout_to_user,workouts RESTART IDENTITY CASCADE")
//	fmt.Println("Test Database migrated")
//	return db
//}
//
//func TestConnect(t *testing.T) {
//	db := setupTestDB()
//	if db == nil {
//		t.Fatal("Failed to connect to the mocked database")
//	}
//	var count int64
//	err := db.Table("users").Count(&count).Error
//	if err != nil {
//		t.Fatalf("migration failed: %v", err)
//	}
//	fmt.Println("PostgreSQL tst database working correctly")
//}
