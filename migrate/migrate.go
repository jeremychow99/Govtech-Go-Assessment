package main

import "example/govtech-test/initializers"
import "example/govtech-test/models"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Teacher{})
	initializers.DB.AutoMigrate(&models.Student{})
}