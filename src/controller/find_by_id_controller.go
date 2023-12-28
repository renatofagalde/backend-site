package controller

import (
	"backend-site/src/config/logger"
	"backend-site/src/view"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func (siteController *siteControllerInterface) FindById(c *gin.Context) {

	logger.Info("controller.FindById", zap.String("journey", "findByID"))
	id := c.Param("id")
	out, _ := json.Marshal(c)
	logger.Info(fmt.Sprintf("context %+v", out), zap.String("journey", "findByID"))

	if reqHeadersBytes, err := json.Marshal(c.Header); err != nil {
		log.Println("Could not Marshal Req Headers")
	} else {
		logger.Info(fmt.Sprintf("header1 %+v", reqHeadersBytes), zap.String("journey", "findByID"))
	}

	// Loop over header names
	for name, values := range c.Request.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
			logger.Info(fmt.Sprintf("header2 %s -> %s", name, value), zap.String("journey", "findByID"))
		}
	}

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
