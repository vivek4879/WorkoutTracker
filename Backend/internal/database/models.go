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
	QueryUserBest(UserId uint, Ex_Id uint) (bestweight float64, reps float64, err error)
	DeleteSession(s Sessions) error
	QueryUserId(userID uint) (*Users, error)
	DeleteUser(u Users) error
	InsertWorkout(UserID uint, workouts []ExerciseData) ([]uint, error)
	InsertWorkoutToUser(userID uint, workoutEntryIDs []uint) error
	UpsertUserBest(userID, exerciseID uint, weight, reps float64) error
	GetAllExercises() ([]Exercises, error)
	GetMeasurements(userID uint) (Measurements, error)
	UpdateMeasurements(userID uint, measurements Measurements) error
	InsertBlankMeasurements(userID uint) error
	GetUserIDByEmail(email string) (uint, error)
	FetchStreakData(userID uint) (*Streak, error)
	UpsertStreak(streakData *Streak) error
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
	User   Users     `gorm:"foreignKey:UserID;reference:ID;constraint:OnDelete:CASCADE"`
}

type Streak struct {
	UserID          uint      `gorm:"column:userid;primaryKey;not null"`
	LastWorkoutDate time.Time `gorm:"not null"`
	CurrentStreak   float64   `gorm:"not null"`
	MaxStreak       float64   `gorm:"not null"`
	User            Users     `gorm:"foreignKey:UserID;reference:ID;constraint:OnDelete:CASCADE"`
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

type Measurements struct {
	Userid     uint     `gorm:"column:userid;primaryKey;not null;foreignKey:UserID;constraint:OnDelete:CASCADE" json:"userid"`
	Weight     *float64 `gorm:"column:weight;" json:"weight,omitempty"`
	Neck       *float64 `gorm:"column:neck;" json:"neck,omitempty"`
	Shoulders  *float64 `gorm:"column:shoulders;" json:"shoulders,omitempty"`
	Chest      *float64 `gorm:"column:chest;" json:"chest,omitempty"`
	LeftBicep  *float64 `gorm:"column:left_bicep;" json:"left_bicep,omitempty"`
	RightBicep *float64 `gorm:"column:right_bicep;" json:"right_bicep,omitempty"`
	UpperAbs   *float64 `gorm:"column:upper_abs;" json:"upper_abs,omitempty"`
	LowerAbs   *float64 `gorm:"column:lower_abs;" json:"lower_abs,omitempty"`
	Waist      *float64 `gorm:"column:waist;" json:"waist,omitempty"`
	Hips       *float64 `gorm:"column:hips;" json:"hips,omitempty"`
	LeftThigh  *float64 `gorm:"column:left_thigh;" json:"left_thigh,omitempty"`
	RightThigh *float64 `gorm:"column:right_thigh;" json:"right_thigh,omitempty"`
	LeftCalf   *float64 `gorm:"column:left_calf;" json:"left_calf,omitempty"`
	RightCalf  *float64 `gorm:"column:right_calf;" json:"right_calf,omitempty"`

	// Foreign Key relationship with the Users table
	User Users `gorm:"foreignKey:Userid;references:ID;constraint:OnDelete:CASCADE" json:"user"` // Foreign Key to Users
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

type Exercises struct {
	ExerciseId       uint   `gorm:"column:exerciseid;primaryKey;autoIncrement"`
	ExerciseName     string `gorm:"column:exercisename"`
	ExerciseImageURL string `gorm:"column:exerciseimageurl"`
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

type UserBests struct {
	UserId            uint      `gorm:"column:userid;primaryKey" json:"userid"`
	Ex_Id             uint      `gorm:"column:ex_id;primaryKey" json:"exerciseid"`
	BestWeight        float64   `json:"bestweight"`
	CorrespondingReps float64   `json:"reps"`
	User              Users     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Exercise          Exercises `gorm:"foreignKey:Ex_Id;references:ExerciseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
