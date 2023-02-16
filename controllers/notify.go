package controllers

import (
	"example/govtech-test/initializers"
	"example/govtech-test/models"
	"fmt"

	// "fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Notify(c *gin.Context) {
	responseCode := 200
	resArr := []string{}
	var body struct {
		Teacher      string `json:"teacher"`
		Notification string `json:"notification"`
	}
	c.Bind(&body)
	// check if teacher valid
	var teacher models.Teacher
	initializers.DB.Preload("teacher").Where("email = ?", body.Teacher).First(&teacher)
	// if teacher in param does not exist in database, throw error response
	if teacher.Email == "" {
		responseCode = 400
		c.JSON(responseCode, gin.H{
			"message": "teacher does not exist",
		})
		return
	}

	// check if there is a valid notification, and if list of students are all valid
	studentArr := strings.Split(body.Notification, " ")[1:]
	notification := strings.Split(body.Notification, " ")[0]
	fmt.Println(studentArr)
	fmt.Println(notification)

	// return list of students who are either registered WITH the teacher , or in the list (AND are NOT SUSPENDED)
	var assignedStudents []models.Student
	initializers.DB.Model(&teacher).Association("AssignedStudents").Find(&assignedStudents)
	for i := range assignedStudents{
		if assignedStudents[i].Suspended == false {
			// then append to res list
			resArr = append(resArr, assignedStudents[i].Email)
		}
	}
	// start by getting list of students registered and not suspended
	for i := range studentArr {
		fmt.Println(studentArr[i])
	}
	// check input students if they exist and are not suspended
	// return valid students in JSON arr
}
