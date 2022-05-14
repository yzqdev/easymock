package controller

import (
	"easymock/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dashboard(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func Wallpaper(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}
