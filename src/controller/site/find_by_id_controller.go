package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (site *siteControllerInterface) FindById(c *gin.Context) {

	fmt.Println("find by id")

	c.JSON(http.StatusOK, "test-gin-adapter-ok")
}
