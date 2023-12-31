package main

import (
	"backend-site/src/config/logger"
	"backend-site/src/controller"
	"backend-site/src/model/repository"
	"backend-site/src/model/service"
	"gorm.io/gorm"
)

func initDependencies(database *gorm.DB) controller.SiteControllerInterface {
	logger.Info("initialize dependencies")

	r := repository.NewSiteReposiroty(database)
	siteDomainService := service.NewSiteDomainService(r)
	return controller.NewSiteControllerInterface(siteDomainService)
}
