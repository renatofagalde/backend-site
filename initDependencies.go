package main

import (
	"backend-site/src/config/logger"
	controller "backend-site/src/controller/site"
	"backend-site/src/model/repository"
	"backend-site/src/model/service"
	"gorm.io/gorm"
)

func initDependencies(database *gorm.DB) controller.SiteControllerInterface {
	logger.Info("initialize dependencies")
	r := repository.NewSiteRepository(database)
	siteDomainService := service.NewSiteDomainService(r)
	return controller.NewSiteControllerInterface(siteDomainService)
}
