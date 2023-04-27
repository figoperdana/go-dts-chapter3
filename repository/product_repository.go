package repository

import (
	"go-jwt/entity"
)

type ProductRepository interface {
	FindById(id string) *entity.Product
}