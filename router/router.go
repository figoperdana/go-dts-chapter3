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
		controllers.CreateAdminUser()
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		// Only admin can perform update and delete operations
		productRouter.PUT("/:productId", middlewares.ProductAuthorization("admin"), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorization("admin"), controllers.DeleteProduct)

		// All users can perform read operations
		productRouter.GET("/:productId", middlewares.ProductAuthorization(""), controllers.GetProduct)
		productRouter.GET("/", middlewares.ProductAuthorization(""), controllers.GetProducts)
	}

	return r
}