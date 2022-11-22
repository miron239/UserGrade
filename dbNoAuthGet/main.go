package main

import (
	"dbAuth/controllers"
	"dbAuth/initializers"
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

	// 2
	r.GET("/userGrades/getAll", controllers.GetAllUserGrades)
	r.GET("/userGrades/getById/:id", controllers.GetUserGradeById)
	r.Run()
}
