package database

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Models struct {
	UserModel MyModel
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

type Workouts struct {
	WorkoutId         uint          `gorm:"column:workoutid;primaryKey;autoIncrement"`
	UserId            uint          `gorm:"column:userid"`
	CurrentExerciseId uint          `gorm:"column:exerciseid"`
	SetNo             uint          `gorm:"column:setno"`
	Repetitions       pq.Int64Array `gorm:"column:repetitions;type:integer[]"`
	Weights           pq.Int64Array `gorm:"column:weights;type:integer[]"`
	CreatedAt         time.Time     `gorm:"column:created_at"`
	User              Users         `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE"`
	Exercise          Exercises     `gorm:"foreignKey:CurrentExerciseId;references:ExerciseId;constraint:OnDelete:CASCADE"`
}

type Exercises struct {
	ExerciseId   uint   `gorm:"column:exerciseid;primaryKey;autoIncrement"`
	ExerciseName string `gorm:"column:exercisename"`
}
