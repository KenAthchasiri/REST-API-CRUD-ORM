package user

import (
	"example/CRUD-REST-API/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBody struct {
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var jsonUser CreateBody

	if err := c.ShouldBindJSON(&jsonUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExist model.User
	model.Db.Where("username = ?", jsonUser.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "User exists"})
		return
	}

	user := model.User{
		Fname:    jsonUser.Fname,
		Lname:    jsonUser.Lname,
		Username: jsonUser.Username,
		Avatar:   jsonUser.Avatar,
	}

	if model.Db.Create(&user); user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "User created",
			"userId":  user.ID,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Create failed",
		})
	}

}
