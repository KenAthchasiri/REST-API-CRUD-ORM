package user

import (
	"example/CRUD-REST-API/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateBody struct {
	Id       uint   `json:"id" binding:"required"`
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func UpdateUser(c *gin.Context) {
	var jsonUser UpdateBody
	var updatedUser model.User

	if err := c.ShouldBindJSON(&jsonUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExist model.User
	model.Db.Where("username = ?", jsonUser.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Username exists"})
		return
	}

	model.Db.First(&updatedUser, jsonUser.Id)
	updatedUser.Fname = jsonUser.Fname
	updatedUser.Lname = jsonUser.Lname
	updatedUser.Username = jsonUser.Username
	updatedUser.Avatar = jsonUser.Avatar

	model.Db.Save(updatedUser)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "User updated",
		"user":    updatedUser,
	})
}
