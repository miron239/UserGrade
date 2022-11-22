package main

import (
	"dbAuth/controllers"
	"dbAuth/initializers"
	"dbAuth/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("hello")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/userGrades/new", middleware.RequireAuth, controllers.CreateNewUserGrade)

	r.PUT("/userGrades/updateById/:id", middleware.RequireAuth, controllers.UpdateUserGrade)
	r.Run()
}
