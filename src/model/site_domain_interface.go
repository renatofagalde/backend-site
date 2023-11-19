package model

func NewSiteDomain(title string) SiteDomainInterface {
	return &siteDomain{
		title: title,
	}
}

type SiteDomainInterface interface {
	GetID() uint
	GetTitle() string
	AtribuirID(uint)
}
