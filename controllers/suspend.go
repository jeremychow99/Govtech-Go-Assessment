package controllers

import (
	// "example/govtech-test/initializers"
	"example/govtech-test/initializers"
	"example/govtech-test/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Suspend(c *gin.Context) {

	var body struct {
		Student string `json:"student"`
	}

	c.Bind(&body)
	fmt.Println(body.Student)

	var dbStudent models.Student
	initializers.DB.Preload("student").Where("email = ?", body.Student).First(&dbStudent)

	fmt.Println(dbStudent.Suspended)
	// check if input student email exists
	if dbStudent.Email == "" {
		fmt.Println("NO EXIST")
		c.JSON(400, gin.H{
			"message": "Student email invalid or does not exist",
		})
	} else {
		if dbStudent.Suspended == false {
			//update suspended status to true
			initializers.DB.Model(&dbStudent).Where("suspended = ?", false).Update("suspended", true)
		} else {
			// if already suspended
			c.JSON(400, gin.H{
				"message": "Student is already suspended",
			})
		}
	}

	c.JSON(204, gin.H{
	})
}
