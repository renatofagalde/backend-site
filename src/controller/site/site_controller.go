package controller

import "github.com/gin-gonic/gin"

type SiteControllerInterface interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
}

type siteControllerInterface struct {
	texto string
}

func NewSiteControllerInterface() SiteControllerInterface {
	return &siteControllerInterface{texto: "Renato"}
}
