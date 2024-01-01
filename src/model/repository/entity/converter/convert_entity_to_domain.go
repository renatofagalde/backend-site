package converter

import (
	"backend-site/src/config/logger"
	"backend-site/src/model"
	"backend-site/src/model/repository/entity"
	"fmt"
	"go.uber.org/zap"
)

func ConvertEntityToDomain(siteEntity entity.SiteEntity) model.SiteDomainInterface {
	domain := model.NewSiteDomain(siteEntity.Title, siteEntity.BannerTitle,
		siteEntity.BannerTitleSlogan, siteEntity.Secao01Title, siteEntity.Secao01TitleDescription)
	domain.AtribuirID(siteEntity.ID)
	logger.Info(
		fmt.Sprintf("convertEntityToDomain ->  %+v", domain), zap.String("journey", "findByID"))

	return domain
}
