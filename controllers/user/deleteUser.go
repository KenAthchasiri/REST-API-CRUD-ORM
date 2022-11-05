package user

import (
	"example/CRUD-REST-API/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	//model.Db.First(&user, id)
	model.Db.First(&user, id).Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "User deleted",
		"user": user,
	})
}
