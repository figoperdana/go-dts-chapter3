package middlewares

import (
	"go-jwt/database"
	"go-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userRole := userData["role"].(string)

		if role == "admin" && userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		} else if role != "admin" && userRole == "admin" {
			c.Next()
			return
		}
		
		Product := models.Product{}
		err = db.Select("user_id").First(&Product, uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data Doesn't Exist",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
