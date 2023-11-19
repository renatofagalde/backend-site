package repository

import (
	"backend-site/src/config/rest_err"
	"backend-site/src/model"
	"gorm.io/gorm"
)

type SiteRepository interface {
	Create(siteDomain model.SiteDomainInterface) (model.SiteDomainInterface, *rest_err.RestErr)
}

type siteRepository struct {
	databaseConnection *gorm.DB
}

func NewSiteReposiroty(database *gorm.DB) SiteRepository {
	return &siteRepository{database}
}
