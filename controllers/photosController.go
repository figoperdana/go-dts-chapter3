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

// CreatePhoto godoc
// @Summary Create photo
// @Description Create photo to post in finalproject
// @Tags photo
// @Accept json
// @Produce json
// @Param title query string true "title"
// @Param caption query string false "caption"
// @Param photo_url query string true "photo_url"
// @Security BearerAuth
// @Success 201 {object} models.Photo "Create photo success"
// @Failure 401 "Unauthorized"
// @Router /photos/create [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// GetAllPhotos godoc
// @Summary Get all photos
// @Description Get all existing photos
// @Tags photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Photo{} "Get all photos success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photos Not Found"
// @Router /photos/getall [get]
func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()
	allPhotos := []models.Photo{}

	db.Find(&allPhotos)

	if len(allPhotos) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No Photos found",
			"error_message": "There are no photos found.",
		})
		return
	}

	c.JSON(http.StatusOK, allPhotos)
}

// GetPhoto godoc
// @Summary Get photo
// @Description Get photo by ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {object} models.Photo{} "Get photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photos/get/{photoId} [get]
func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	Photo.ID = uint(photoId)

	err := db.First(&Photo, "id = ?", photoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// UpdatePhoto godoc
// @Summary Update photo
// @Description Update photo identified by given ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {object} models.Photo{} "Update photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photos/update/{photoId} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete photo identified by given ID
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {string} string "Delete photo success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /photos/delete/{photoId} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Delete Success",
		"Message": "The photo has been successfully deleted",
	})
}
