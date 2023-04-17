package controllers

import (
	"finalproject/database"
	"finalproject/helpers"
	"finalproject/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateSocialMedia godoc
// @Summary Create social media
// @Description Create social media of the user
// @Tags social media
// @Accept json
// @Produce json
// @Param name query string true "name"
// @Param social_media_url query string true "social_media_url"
// @Security BearerAuth
// @Success 201 {object} models.SocialMedia "Create social media success"
// @Failure 401 "Unauthorized"
// @Router /socialmedia/create [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

// GetAllSocialMedia godoc
// @Summary Get all social media
// @Description Get all social media in finalproject
// @Tags social media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.SocialMedia "Get all social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/getall [get]
func GetAllSocialMedias(c *gin.Context) {
	db := database.GetDB()
	allSocialMedias := []models.SocialMedia{}

	db.Find(&allSocialMedias)

	if len(allSocialMedias) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No social media found",
			"error_message": "There are no social media found.",
		})
		return
	}

	c.JSON(http.StatusOK, allSocialMedias)
}

// GetSocialMedia godoc
// @Summary Get social media
// @Description Get social media identified by given id
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {object} models.SocialMedia "Get social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/get/{socialMediaId} [get]
func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	SocialMedia.ID = uint(socialMediaId)

	err := db.First(&SocialMedia, "id = ?", socialMediaId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update social media
// @Description Update social media identified by given id
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {object} models.SocialMedia "Update social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/update/{socialMediaId} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete social media
// @Description Delete social media identified by given ID
// @Tags social media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "ID of the social media"
// @Security BearerAuth
// @Success 200 {string} string "Delete social media success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Social Media Not Found"
// @Router /socialmedia/delete/{socialMediaId} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	err := db.Where("id = ?", socialMediaId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Delete Success",
		"Message": "The social media has been successfully deleted",
	})
}
