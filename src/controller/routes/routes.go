package routes

import (
	controller "backend-site/src/controller/site"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/", controller.FindById)
	r.POST("/", controller.Create)
}
