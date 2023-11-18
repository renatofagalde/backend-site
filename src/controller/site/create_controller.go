package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (siteController *siteControllerInterface) Create(c *gin.Context) {

	fmt.Println("create")
	var text = fmt.Sprintf("test-gin-adapter-ok  -> %s", siteController.texto)
	c.JSON(http.StatusOK, text)
}
