package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Email string `gorm:"unique"`
	AssignedStudents []Student `gorm:"many2many:assigned_students;"`
}
