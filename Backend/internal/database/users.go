package database

import (
	"fmt"
	"time"
)

func (u MyModel) InsertSession(Id uint, Token string, expiry time.Time) error {
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
