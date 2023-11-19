package view

import (
	"backend-site/src/controller/model/response"
	"backend-site/src/model"
)

func ConvertDomainToResponse(domain model.SiteDomainInterface) response.SiteResponse {
	return response.SiteResponse{
		Title: domain.GetTitle(),
	}
}
