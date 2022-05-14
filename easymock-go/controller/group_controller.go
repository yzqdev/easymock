package controller

import (
	"easymock/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GroupGetList(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func GroupJoin(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func GroupCreate(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func GroupUpdate(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}

func GroupDelete(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "aaa")
}
