package controller

import "github.com/gin-gonic/gin"

type SiteControllerInterface interface {
	Create(c *gin.Context)
}
