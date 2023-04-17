package controllers

import (
	"github.com/gin-gonic/gin"
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
)

var (
	appJSON = "application/json"
)

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"full_name": user.FullName,
		"email":     user.Email,
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	hashedPassword := helpers.HashPass(password)
	if User.Password != hashedPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
