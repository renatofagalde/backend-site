package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (site *siteControllerInterface) Create(c *gin.Context) {

	fmt.Println("create")
	c.JSON(http.StatusOK, "test-gin-adapter-ok")
}
