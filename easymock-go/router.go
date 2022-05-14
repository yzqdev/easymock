package main

import (
	"easymock/controller"
	"easymock/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	e.POST("/dashboard", controller.Dashboard)
	e.POST("/wallpaper", controller.Wallpaper)
	baseRouter := e.Group("/auth")
	{
		baseRouter.POST("/login", controller.Login)
		baseRouter.POST("/reg", controller.Register)
		baseRouter.POST("/update", middleware.JwtHandler(), controller.Register)
		baseRouter.POST("/outLogin")
	}
	projectRouter := e.Group("/project", middleware.JwtHandler())
	{

		projectRouter.GET("/index", controller.ProjectGetList)
		projectRouter.POST("/copy", controller.ProjectCopy)
		projectRouter.POST("/create", controller.ProjectCreate)
		projectRouter.POST("/update", controller.ProjectUpdate)
		projectRouter.POST("/sync/swagger", controller.ProjectUpdateSwagger)
		projectRouter.POST("/update_workbench", controller.ProjectUpdateWorkbench)
		projectRouter.POST("/delete", controller.ProjectDelete)

	}
	mockRouter := e.Group("/mock", middleware.JwtHandler())
	{
		mockRouter.GET("/index", controller.MockGetList)
		mockRouter.GET("/create", controller.MockCreate)
		mockRouter.GET("/update", controller.MockUpdate)
		mockRouter.GET("/delete", controller.MockDelete)
		mockRouter.GET("/export", controller.MockExport) //todo 看前端

	}
	groupRouter := e.Group("/group", middleware.JwtHandler())
	{
		groupRouter.POST("/join", controller.GroupJoin)
		groupRouter.POST("/create", controller.GroupCreate)
		groupRouter.POST("/update", controller.GroupUpdate)
		groupRouter.POST("/delete", controller.GroupDelete)
	}
}
