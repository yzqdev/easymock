package controller

import (
	"easymock/utils"
	"github.com/gin-gonic/gin"
)

func MockGetList(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func MockCreate(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func MockUpdate(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func MockDelete(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func MockExport(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}
