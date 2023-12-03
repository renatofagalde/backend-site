package routes

import (
	controller "backend-site/src/controller/site"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, siteController controller.SiteControllerInterface) {

	r.GET("/:id", siteController.FindById)
	r.POST("/", siteController.Create)
}
