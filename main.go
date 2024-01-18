package main

import (
	"github.com/gin-gonic/gin"
	"plume-backend/controllers"
	"plume-backend/initilializers"
	"plume-backend/midlleware"
)

func init() {
	initilializers.LoadEnvVariables()
	initilializers.ConnectToDB()
	initilializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", midlleware.RequireAuth, controllers.Validate)
	r.Run()
}
