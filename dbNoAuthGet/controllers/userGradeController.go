package controllers

import (
	"dbAuth/initializers"
	userModel "dbAuth/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUserGrades(c *gin.Context) {
	//get the post
	var userGrades []userModel.UserGrade
	initializers.DB.Find(&userGrades)
	c.JSON(http.StatusOK, gin.H{
		"userGrades": userGrades,
	})
}

func GetUserGradeById(c *gin.Context) {
	//get the post
	id := c.Param("id")
	var userGrade userModel.UserGrade
	initializers.DB.First(&userGrade, id)
	c.JSON(http.StatusOK, gin.H{
		"userGrade": userGrade,
	})
}
