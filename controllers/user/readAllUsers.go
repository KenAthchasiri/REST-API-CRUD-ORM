package user

import (
	"example/CRUD-REST-API/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	var users []model.User
	model.Db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "User read success",
		"userId":  users,
	})
}
