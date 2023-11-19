package model

func NewSiteDomain(title string) SiteDomainInterface {
	return &siteDomain{
		title: title,
	}
}

type SiteDomainInterface interface {
	GetID() string
	GetTitle() string
	AtribuirID(uint)
}
