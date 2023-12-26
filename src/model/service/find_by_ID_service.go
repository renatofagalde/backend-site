package service

import (
	"backend-site/src/config/logger"
	"backend-site/src/config/rest_err"
	"backend-site/src/model"
	"go.uber.org/zap"
)

func (siteDomainService *siteDomainService) FindByID(id string) (model.SiteDomainInterface, *rest_err.RestErr) {
	logger.Info("service.FindByID", zap.String("journey", "findByID"))
	return siteDomainService.siteRepository.FindByID(id)
}
