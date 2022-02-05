package controller

import (
	"easymock/utils"
	"github.com/gin-gonic/gin"
)

func GroupGetList(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func GroupJoin(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func GroupCreate(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func GroupUpdate(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func GroupDelete(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}
