package service

import (
	"backend-site/src/config/logger"
	"backend-site/src/config/rest_err"
	"backend-site/src/model"
)

func (s *siteDomainService) Create(domainInterface model.SiteDomainInterface) (model.SiteDomainInterface, *rest_err.RestErr) {
	logger.Info("init create model")

	siteRepository, err := s.siteRepository.Create(domainInterface)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	return siteRepository, nil
}
