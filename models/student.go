package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Email            string `gorm:"unique"`
	Suspended        bool
	AssignedTeachers []Teacher `gorm:"many2many:assigned_students;"`
}
