package controller

import (
	"easymock/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProjectGetList(c *gin.Context) {

	utils.JSON(c, http.StatusOK, "chengg", "111")
}

func ProjectCopy(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "count")
}

func ProjectCreate(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "count")
}

func ProjectUpdate(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "count")
}

func ProjectUpdateSwagger(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "count")
}

func ProjectUpdateWorkbench(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "count")
}

func ProjectDelete(c *gin.Context) {
	utils.JSON(c, http.StatusOK, "success", "count")
}
