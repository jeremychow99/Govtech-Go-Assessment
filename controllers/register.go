package controllers

import (
	"example/govtech-test/initializers"
	"example/govtech-test/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	message := ""
	responseCode := 204
	responseStudents := []string{}
	var body struct {
		Teacher  string   `json:"teacher" binding:"required"`
		Students []string `json:"students" binding:"required"`
	}

	c.Bind(&body)

	// create teacher
	var dbTeacher models.Teacher

	initializers.DB.Preload("teacher").Where("email = ?", body.Teacher).First(&dbTeacher)
	teacher := models.Teacher{Email: body.Teacher, AssignedStudents: []models.Student{}}

	if dbTeacher.Email == "" {
		initializers.DB.Create(&teacher)
	} else {
		teacher = dbTeacher
	}

	// create students and associations
	var assignedStudents []models.Student
	initializers.DB.Model(&teacher).Association("AssignedStudents").Find(&assignedStudents)
	for i := range body.Students {
		initializers.DB.Model(&teacher).Association("AssignedStudents")

		student := models.Student{Email: body.Students[i], Suspended: false, AssignedTeachers: []models.Teacher{}}
		var dbStudent models.Student
		initializers.DB.Preload("student").Where("email = ?", student.Email).First(&dbStudent)

		// if student does not exist in database, create student entry
		if dbStudent.Email == "" {
			initializers.DB.Create(&student)
		} else {
			// maintain suspension status
			student.Suspended = dbStudent.Suspended
		}
		// check for association
		for i := range assignedStudents {
			if assignedStudents[i].Email == student.Email {
				responseCode = 409
				responseStudents = append(responseStudents, student.Email)
				message = "Some students are already registered (check 'failed_to_register'). Students not listed in the 'failed_to_register' list have been registered."
				break
			}
		}

		initializers.DB.Model(&teacher).Association("AssignedStudents").Append([]*models.Student{&student})
		initializers.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&teacher)
	}

	// response
	c.JSON(responseCode, gin.H{
		"message":            message,
		"failed_to_register": responseStudents,
	})

}
