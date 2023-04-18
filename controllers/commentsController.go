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

// CreateComments godoc
// @Summary Create comments
// @Description Create comments for photo identified by given id
// @Tags comments
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Param message query string true "message"
// @Security BearerAuth
// @Success 201 {object} models.Comment "Create comments success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Photo Not Found"
// @Router /comments/create/{photoId} [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	Photo := models.Photo{}

	err := db.Select("user_id").First(&Photo, uint(photoId)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Photo Not Found",
			"message": "Photo doesn't exist, failed to create comments",
		})
		return
	}

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = uint(photoId)

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// GetAllComments godoc
// @Summary Get all comments
// @Description Get all comments in finalproject
// @Tags comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Comment "Get all comments success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comments Not Found"
// @Router /comments/getall [get]
func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	allComments := []models.Comment{}

	db.Find(&allComments)

	if len(allComments) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No comments found",
			"error_message": "There are no comments found.",
		})
		return
	}

	c.JSON(http.StatusOK, allComments)
}

// GetAllCommentsForPhoto godoc
// @Summary Get all comments for specific photo
// @Description Get all comments for photo with given id
// @Tags comment
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo"
// @Security BearerAuth
// @Success 200 {object} []models.Comment "Get all comments success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comments Not Found"
// @Router /comments/getall/{photoId} [get]
func GetAllCommentsForPhoto(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	allComments := []models.Comment{}

	db.Where("photo_id = ?", photoId).Find(&allComments)

	if len(allComments) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No Comments found",
			"error_message": "There are no comments found for this photo.",
		})
		return
	}

	c.JSON(http.StatusOK, allComments)
}

// GetComments godoc
// @Summary Get comments
// @Description Get comments identified by given id
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comments"
// @Security BearerAuth
// @Success 200 {object} models.Comment "Get comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comment Not Found"
// @Router /comments/get/{commentId} [get]
func GetComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment.ID = uint(commentId)

	err := db.First(&Comment, "id = ?", commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// UpdateComment godoc
// @Summary Update comment
// @Description Update comment identified by given id
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Security BearerAuth
// @Success 200 {object} models.Comment "Update comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comment Not Found"
// @Router /comments/update/{commentId} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.First(&Comment, "id = ?", commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment identified by given ID
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment"
// @Security BearerAuth
// @Success 200 {string} string "Delete comment success"
// @Failure 401 "Unauthorized"
// @Failure 404 "Comment Not Found"
// @Router /comments/delete/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Delete Success",
		"Message": "The comment has been successfully deleted",
	})
}
