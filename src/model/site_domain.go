package model

type siteDomain struct {
	id    uint
	title string
}

func (sd *siteDomain) GetID() uint {
	return sd.id
}

func (sd *siteDomain) GetTitle() string {
	return sd.title
}

func (sd *siteDomain) AtribuirID(id uint) {
	sd.id = id
}
