package converter

import (
	"backend-site/src/model"
	"backend-site/src/model/repository/entity"
)

func ConvertEntityToDomain(siteEntity entity.SiteEntity) model.SiteDomainInterface {
	domain := model.NewSiteDomain(siteEntity.Title)
	domain.AtribuirID(siteEntity.ID)
	return domain
}
