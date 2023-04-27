package service

import (
	"go-jwt/entity"
	"go-jwt/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}


func setupTest(t *testing.T) {
	productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	productService = ProductService{Repository: productRepository}
	t.Cleanup(func() {
		productRepository.Mock = mock.Mock{}
	})
}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	setupTest(t)

	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")

}

func TestProductServiceGetOneProduct(t *testing.T) {
	setupTest(t)

	product := entity.Product{
		Id:          "1",
		Title:       "Telur Mentah",
		Description: "Dikeluarkan oleh ayam terpilih",
	}

	productRepository.Mock.On("FindById", "1").Return(product)

	result, err := productService.GetOneProduct("1")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be '1'")
	assert.Equal(t, product.Title, result.Title, "result has to be 'Telur Mentah'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'Dikeluarkan oleh ayam terpilih'")
}