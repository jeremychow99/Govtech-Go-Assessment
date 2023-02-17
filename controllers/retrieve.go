package controllers

import (
	"example/govtech-test/initializers"
	"example/govtech-test/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// helper function to get common students
func intersection(arr1, arr2 []string) []string {
	out := []string{}
	bucket := map[string]bool{}
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}
	return out
}

func Retrieve(c *gin.Context) {
	responseCode := 200
	responseArr := []string{}
	// get query string
	query := c.Request.URL.Query()["teacher"]

	var assignedStudents []models.Student

	for index := range query {
		var teacher models.Teacher
		initializers.DB.Model(&teacher).Association("AssignedStudents")
		initializers.DB.Preload("teacher").Where("email = ?", query[index]).First(&teacher)
		// if teacher in param does not exist in database, throw error response
		if teacher.Email == "" {
			responseCode = 400
			c.JSON(responseCode, gin.H{
				"message": "one or more teachers specified do not exist",
			})
			return
		}

		initializers.DB.Model(&teacher).Association("AssignedStudents").Find(&assignedStudents)
		testArr := []string{}
		if index == 0 {
			for i := range assignedStudents {
				fmt.Println(i)
				fmt.Println(assignedStudents[i])
				responseArr = append(responseArr, assignedStudents[i].Email)
			}

			fmt.Println(responseArr)
		} else {

			for i := range assignedStudents {
				fmt.Println(i)
				fmt.Println(assignedStudents[i])
				testArr = append(testArr, assignedStudents[i].Email)
			}
			fmt.Println("ARRAYS AS FOLLOWS")
			fmt.Println(testArr)
			fmt.Println(responseArr)
			responseArr = intersection(responseArr, testArr)

		}
	}

	c.JSON(responseCode, gin.H{
		"students": responseArr,
	})
}
