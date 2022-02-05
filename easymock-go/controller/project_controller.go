package controller

import (
	"easymock/utils"
	"github.com/gin-gonic/gin"
)

func ProjectGetList(c *gin.Context) {

	utils.JSON(c, 200, "chengg", "111")
}

func ProjectCopy(c *gin.Context) {
	utils.JSON(c, 200, "success", "count")
}

func ProjectCreate(c *gin.Context) {
	utils.JSON(c, 200, "success", "count")
}

func ProjectUpdate(c *gin.Context) {
	utils.JSON(c, 200, "success", "count")
}

func ProjectUpdateSwagger(c *gin.Context) {
	utils.JSON(c, 200, "success", "count")
}

func ProjectUpdateWorkbench(c *gin.Context) {
	utils.JSON(c, 200, "success", "count")
}

func ProjectDelete(c *gin.Context) {
	utils.JSON(c, 200, "success", "count")
}
