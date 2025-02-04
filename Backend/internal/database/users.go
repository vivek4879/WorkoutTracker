package database

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

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

func (u MyModel) INSERTSESSION(Id uint, Token string, expiry time.Time) error {
	session := Sessions{
		UserID: Id,
		Token:  Token,
		Expiry: expiry,
	}

	res := u.db.Create(&session)
	if res.Error != nil {
		fmt.Println("Error inserting new session", res.Error)
		return res.Error
	}
	return nil
}
func (u MyModel) QUERYSESSION(SessionToken string) (*Sessions, error) {
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

func (u MyModel) Query(Email string) (*UserLogin, error) {
	var user UserLogin
	res := u.db.Table("users").Select("email,password,userid").Where("email = ?", Email).First(&user)

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
