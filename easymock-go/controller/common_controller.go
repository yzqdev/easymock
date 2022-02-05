package controller

import (
	"easymock/utils"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}

func Wallpaper(c *gin.Context) {
	utils.JSON(c, 200, "success", "aaa")
}
