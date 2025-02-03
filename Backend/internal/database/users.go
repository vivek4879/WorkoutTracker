package database

import (
	"fmt"
	"gorm.io/gorm"
)

type MyModel struct {
	db *gorm.DB
}

type Users struct {
	FirstName string `gorm:"column:firstname"`
	LastName  string `gorm:"column:lastname"`
	Email     string
	Password  string
}

type UserLogin struct {
	Email    string
	Password string
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
	res := u.db.Table("users").Select("email,password").Where("email = ?", Email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
