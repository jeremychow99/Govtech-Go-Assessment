package controllers

import (
	"example/govtech-test/initializers"
	"example/govtech-test/models"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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

	// NEED TO CHECK IF EXISTS: MAYBE USING models.Teacher.Email 
	fmt.Println("===============")
	fmt.Println(dbTeacher.Email) // type of models.teacher
	fmt.Println(reflect.TypeOf(dbTeacher.Email))
	fmt.Println("===============")
	if dbTeacher.Email == ""{
		initializers.DB.Create(&teacher)
	} else {
		teacher = dbTeacher
	}

	// create students and associations
	for i := range body.Students {
		initializers.DB.Model(&teacher).Association("AssignedStudents")
		student := models.Student{Email: body.Students[i], Suspended: false, AssignedTeachers: []models.Teacher{}}
		initializers.DB.Create(&student)
		result := initializers.DB.Model(&teacher).Association("AssignedStudents").Append([]*models.Student{&student});

		// check for duplicate error, change response code and message accordingly
		_, ok := result.(*mysql.MySQLError)
		if result != nil && ok && result.(*mysql.MySQLError).Number == 1452 {
			responseStudents = append(responseStudents, student.Email)
			message = "failed to register some students, they are already registered"
			responseCode = 409
		  }
		  
		initializers.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&teacher)
	}

	// response
	c.JSON(responseCode, gin.H{
		"message": message,
		"failed_to_register": responseStudents,
	})

}
