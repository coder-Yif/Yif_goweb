package main

import (
	"awesomeProject3/controller"
	"awesomeProject3/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/Info", middleware.AuthMiddleWare(), controller.Info)
	return r
}
