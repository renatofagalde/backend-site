package controller

import (
	"backend-site/src/config/logger"
	"backend-site/src/view"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (siteController *siteControllerInterface) FindById(c *gin.Context) {

	logger.Info("controller.FindById", zap.String("journey", "findByID"))
	id := c.Param("id")
	out, _ := json.Marshal(c)
	logger.Info(fmt.Sprintf("context %+v", out), zap.String("journey", "findByID"))

	logger.Info(fmt.Sprintf("context %+v", out), zap.String("journey", "findByID"))
	logger.Info(fmt.Sprintf("request header %+v", c.Request.Header), zap.String("journey", "findByID"))
	logger.Info(fmt.Sprintf("request header %+v", c.Request.Host), zap.String("journey", "findByID"))

	siteDomain, err := siteController.service.FindByID(id)
	if err != nil {
		c.JSON(err.Code, err)
		message := "Error retriving by ID"
		logger.Error(message, err)
		return
	}
	result := view.ConvertDomainToResponse(siteDomain)
	logger.Info(fmt.Sprintf("init FindUserByID find_controller successfuly %+v", result), zap.String("journey", "findByID"))
	c.JSON(http.StatusOK, result)
}
