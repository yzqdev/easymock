package controller

import (
	"easymock/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MockGetList(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func MockCreate(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func MockUpdate(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func MockDelete(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func MockExport(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}
