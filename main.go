package main

import (
	UserControllers "example/CRUD-REST-API/controllers/user"
	"example/CRUD-REST-API/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDb()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/users", UserControllers.ReadAll)
	r.GET("/user/:id", UserControllers.OneUser)
	r.POST("/users", UserControllers.CreateUser)
	r.PUT("/users",UserControllers.UpdateUser)
	r.DELETE("/user/:id", UserControllers.DeleteUser)
	r.Run("localhost:8080")
}
