package routes

import (
	"backend-site/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, siteController controller.SiteControllerInterface) {

	r.GET("/site/:id", siteController.FindById)
	r.POST("/site/", siteController.Create)
}
