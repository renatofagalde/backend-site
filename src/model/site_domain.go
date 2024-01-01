package model

type siteDomain struct {
	id                      uint
	title                   string
	bannerTitle             string
	bannerTitleSlogan       string
	secao01Title            string
	secao01TitleDescription string
}

func (sd *siteDomain) GetID() uint {
	return sd.id
}

func (sd *siteDomain) GetTitle() string {
	return sd.title
}

func (sd *siteDomain) GetBannerTitle() string {
	return sd.bannerTitle
}

func (sd *siteDomain) GetBannerTitleSlogan() string {
	return sd.bannerTitleSlogan
}

func (sd *siteDomain) GetSecao01Title() string {
	return sd.secao01Title
}
func (sd *siteDomain) GetSecao01TitleDescription() string {
	return sd.secao01TitleDescription
}

func (sd *siteDomain) AtribuirID(id uint) {
	sd.id = id
}
