package controllers

import (
	"dbAuth/initializers"
	userModel "dbAuth/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNewUserGrade(c *gin.Context) {
	var body struct {
		UserId        string
		PostpaidLimit int
		Spp           int
		ShippingFee   int
		ReturnFee     int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	UserGrade := userModel.UserGrade{UserId: body.UserId, PostpaidLimit: body.PostpaidLimit, Spp: body.Spp,
		ShippingFee: body.ShippingFee, ReturnFee: body.ReturnFee}

	result := initializers.DB.Create(&UserGrade) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create userGrade",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

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

func UpdateUserGrade(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		UserId        string
		PostpaidLimit int
		Spp           int
		ShippingFee   int
		ReturnFee     int
	}

	c.Bind(&body)
	var userGrade userModel.UserGrade
	initializers.DB.First(&userGrade, id)
	initializers.DB.Model(&userGrade).Updates(userModel.UserGrade{UserId: body.UserId, PostpaidLimit: body.PostpaidLimit,
		Spp: body.Spp, ShippingFee: body.ShippingFee, ReturnFee: body.ReturnFee})

	c.JSON(http.StatusOK, gin.H{
		"userGrades": userGrade,
	})
}
