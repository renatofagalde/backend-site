package repository

import (
	"backend-site/src/config/rest_err"
	"backend-site/src/model"
	"backend-site/src/model/repository/entity"
	"backend-site/src/model/repository/entity/converter"
)

func (siteRepository *siteRepository) FindByID(id string) (model.SiteDomainInterface, *rest_err.RestErr) {

	siteEntity := entity.SiteEntity{}

	siteRepository.databaseConnection.
		First(&siteEntity, id)

	return converter.ConvertEntityToDomain(siteEntity), nil
}
