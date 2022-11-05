package user

import (
	"example/CRUD-REST-API/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OneUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	model.Db.Find(&user, id)

	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "User read success",
			"userId":  user,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "User not exists",
		})
	}

}
