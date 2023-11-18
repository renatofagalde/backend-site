package controller

import "github.com/gin-gonic/gin"

type SiteControllerInterface interface {
	Create(c *gin.Context)
}

//
//type siteControllerInterface struct {
//
//}
//
//func NewSiteControllerInterface(texto string) SiteControllerInterface {
//	return &siteControllerInterface{string: texto}
//}
