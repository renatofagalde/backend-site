package controller

import (
	"backend-site/src/config/logger"
	"backend-site/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (siteController *siteControllerInterface) FindById(c *gin.Context) {

	logger.Info("controller.FindById", zap.String("journey", "findByID"))
	id := c.Param("id")

	siteDomain, err := siteController.service.FindByID(id)
	if err != nil {
		c.JSON(err.Code, err)
		message := "Error retriving by ID"
		logger.Error(message, err)
		return
	}
	logger.Info("init FindUserByID find_controller successfuly")
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(siteDomain))
}
