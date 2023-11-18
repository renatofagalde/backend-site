package controller

import "github.com/gin-gonic/gin"

type SiteControllerInterface interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
}

type siteControllerInterface struct {
	texto string
}

func NewSiteControllerInterface(texto string) SiteControllerInterface {
	return &siteControllerInterface{texto: texto}
}
