package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindById(c *gin.Context) {

	fmt.Println("dentro do FindById")

	c.JSON(http.StatusOK, "test-gin-adapter-ok")
}
