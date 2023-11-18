package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (siteController *siteControllerInterface) FindById(c *gin.Context) {

	c.JSON(http.StatusOK, "test-gin-adapter-ok")
}
