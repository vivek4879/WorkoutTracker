package database

import (
	"gorm.io/gorm"
)

type Models struct {
	UserModel MyModel
}

func NewModels(conn *gorm.DB) Models {
	return Models{UserModel: MyModel{conn}}
}
