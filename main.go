package main

import (
	"example/govtech-test/controllers"
	"example/govtech-test/initializers"

	"github.com/gin-gonic/gin"
)

// runs before main()
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/api/register", controllers.Register)
	r.GET("/api/commonstudents", controllers.Retrieve)
	r.POST("/api/suspend", controllers.Suspend)
	r.POST("/api/retrievefornotifications",controllers.Notify)

	r.Run()
}
