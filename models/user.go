package models

import (
	"go-jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" validate:"required, fullname"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" validate:"required, email"`
	Password string    `gorm:"not null" json:"password"  validate:"required, password"`
	Products []Product `json:"products"`
	Role     string    `gorm:"default:user" json:"role" validate:"required-Role is required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}