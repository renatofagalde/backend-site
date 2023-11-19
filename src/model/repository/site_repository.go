package repository

import (
	"backend-site/src/config/rest_err"
	"backend-site/src/model"
	"gorm.io/gorm"
)

type siteRepository struct {
	databaseConnection *gorm.DB
}

func NewSiteRepository(database *gorm.DB) SiteRepository {
	return &siteRepository{databaseConnection: database}
}

type SiteRepository interface {
	Create(siteDomain model.SiteDomainInterface) (model.SiteDomainInterface, *rest_err.RestErr)
}
