package entity

import "gorm.io/gorm"

type SiteEntity struct {
	gorm.Model
	TenentID                string
	Title                   string
	UsuarioID               string
	BannerTitle             string
	BannerTitleSlogan       string
	Secao01Title            string
	Secao01TitleDescription string
}
