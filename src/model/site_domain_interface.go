package model

func NewSiteDomain(title string, bannerTitle string, bannerTitleSlogan string,
	secao01Title string, secao01TitleDescription string) SiteDomainInterface {
	return &siteDomain{
		title:                   title,
		bannerTitle:             bannerTitle,
		bannerTitleSlogan:       bannerTitleSlogan,
		secao01Title:            secao01Title,
		secao01TitleDescription: secao01TitleDescription,
	}
}

type SiteDomainInterface interface {
	GetID() uint
	GetTitle() string
	AtribuirID(uint)
}
