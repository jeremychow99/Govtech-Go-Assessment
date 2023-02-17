package controllers

import (
	"example/govtech-test/initializers"
	"example/govtech-test/models"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)
// helper function to remove duplicates
func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
	   if _, ok := bucket[str]; !ok {
		  bucket[str] = true
		  result = append(result, str)
	   }
	}
	return result
 }

func Notify(c *gin.Context) {
	responseCode := 200
	resArr := []string{}
	var body struct {
		Teacher      string `json:"teacher"`
		Notification string `json:"notification"`
	}
	c.Bind(&body)

	// check if teacher valid, exists in DB
	var teacher models.Teacher
	initializers.DB.Preload("teacher").Where("email = ?", body.Teacher).First(&teacher)
	// if teacher does not exist in database, throw error response
	if teacher.Email == "" {
		responseCode = 400
		c.JSON(responseCode, gin.H{
			"message": "teacher does not exist",
		})
		return
	}

	// get arr of students
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

	// for students not registered with the teacher
	for i := range studentArr {
		// check if start with @
		if studentArr[i][0:1] == "@" {
			var dbStudent models.Student
			fmt.Println(studentArr[i])
			initializers.DB.Preload("student").Where("email = ?", studentArr[i][1:]).First(&dbStudent)

			// then check if student exists and not suspended
			if dbStudent.Email != "" && dbStudent.Suspended == false {
				resArr = append(resArr, dbStudent.Email)
			}

		}
	}
	// return valid students in JSON array
	c.JSON(responseCode, gin.H{
		"recipients": removeDuplicates(resArr),
	})
}
