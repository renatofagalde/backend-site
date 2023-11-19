package service

import (
	"backend-site/src/config/rest_err"
	"backend-site/src/model"
	"backend-site/src/model/repository"
)

func NewSiteDomainService(siteRepository repository.SiteRepository) SiteDomainService {
	return &siteDomainService{siteRepository}
}

type siteDomainService struct {
	siteRepository repository.SiteRepository
}

type SiteDomainService interface {
	Create(domainInterface model.SiteDomainInterface) (model.SiteDomainInterface, *rest_err.RestErr)
}
