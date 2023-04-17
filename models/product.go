package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"  validate:"required, title"`
	Description string `json:"description"  validate:"required, description"`
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}