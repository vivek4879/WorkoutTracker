package database

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

// newInterface
type UserModelInterface interface {
	Insert(firstName, lastName, email, password string) error
	Query(email string) (*Users, error)
	InsertSession(Id uint, Token string, expiry time.Time) error
	QuerySession(SessionToken string) (*Sessions, error)

	DeleteSession(s Sessions) error
	QueryUserId(userID uint) (*Users, error)
	DeleteUser(u Users) error
	InsertWorkout(UserID uint, workouts []ExerciseData) ([]uint, error)
	InsertWorkoutToUser(userID uint, workoutEntryIDs []uint) error
}

// Ensure MyModel implements UserModelInterface
var _ UserModelInterface = (*MyModel)(nil)

// EndnewInterface
type Models struct {
	UserModel UserModelInterface
}

func NewMyModel(db *gorm.DB) UserModelInterface {
	return &MyModel{db: db}
}
func NewModels(conn *gorm.DB) Models {
	return Models{UserModel: MyModel{conn}}
}

type MyModel struct {
	db *gorm.DB
}
type Sessions struct {
	UserID uint      `gorm:"column:userid;primaryKey;not null"`
	Token  string    `gorm:"unique;not null"`
	Expiry time.Time `gorm:"not null"`
	User   Users     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Users struct {
	ID        uint   `gorm:"column:userid;primaryKey;autoIncrement"` // Use `ID` (uppercase) for GORM
	FirstName string `gorm:"column:firstname;not null"`
	LastName  string `gorm:"column:lastname;not null"`
	Email     string `gorm:"column:email;unique;not null"`
	Password  string `gorm:"column:password; not null"`
}

type UserLogin struct {
	Email    string
	Password string
	Userid   uint
}

type WorkoutToUser struct {
	UserID                  uint     `gorm:"column:userid;primaryKey"`
	WorkoutEntryIDSecondary uint     `gorm:"column:workout_entry_id_Secondary;primaryKey;"`
	Workout                 Workouts `gorm:"foreignKey:WorkoutEntryIDSecondary;references:WorkoutEntryID;constraint:OnDelete:CASCADE"`
}
type Workouts struct {
	WorkoutEntryID    uint          `gorm:"column:workout_entry_id;primaryKey;autoIncrement"`
	WorkoutId         uint          `gorm:"column:workoutid"`
	UserId            uint          `gorm:"column:userid"`
	CurrentExerciseId uint          `gorm:"column:exerciseid"`
	SetNo             pq.Int64Array `gorm:"column:setno;type:integer[]"`
	Repetitions       pq.Int64Array `gorm:"column:repetitions;type:integer[]"`
	Weights           pq.Int64Array `gorm:"column:weights;type:integer[]"`
	CreatedAt         time.Time     `gorm:"column:created_at"`
	User              Users         `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE"`
	Exercise          Exercises     `gorm:"foreignKey:CurrentExerciseId;references:ExerciseId;constraint:OnDelete:CASCADE"`
}

//func (Workouts) TableName() string {
//	return "workouts"
//}

type Exercises struct {
	ExerciseId   uint   `gorm:"column:exerciseid;primaryKey;autoIncrement"`
	ExerciseName string `gorm:"column:exercisename"`
}

// struct to capture frontend data
type WorkoutSet struct {
	SetNo       uint  `json:"setno"`
	Repetitions int64 `json:"repetitions"`
	Weight      int64 `json:"weights"`
}

// struct to capture frontend data
type WorkoutInput struct {
	UserId   uint           `json:"userid"`
	Workouts []ExerciseData `json:"workouts"`
}

// struct to capture frontend data
type ExerciseData struct {
	ExerciseId uint         `json:"exerciseid"`
	Sets       []WorkoutSet `json:"sets"`
	CreatedAt  time.Time    `json:"created_at"`
}
