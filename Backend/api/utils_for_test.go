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
