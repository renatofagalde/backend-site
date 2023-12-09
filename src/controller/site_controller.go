package controller

import (
	"backend-site/src/model/service"
	"github.com/gin-gonic/gin"
)

type SiteControllerInterface interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
}

type siteControllerInterface struct {
	service service.SiteDomainService
}

func NewSiteControllerInterface(serviceInterface service.SiteDomainService) SiteControllerInterface {
	return &siteControllerInterface{service: serviceInterface}
}
