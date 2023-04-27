package router

import (
	"go-jwt/controllers"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)

	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.GET("/", controllers.GetAllProducts)

		productRouter.GET("/:productId", middlewares.ProductAuthorization("GET"), controllers.GetProduct)

		productRouter.PUT("/:productId", middlewares.ProductAuthorization("UPDATE"), controllers.UpdateProduct)

		productRouter.DELETE("/:productId", middlewares.ProductAuthorization("DELETE"), controllers.DeleteProduct)
	}

	return r
}