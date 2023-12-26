package entity

import "gorm.io/gorm"

type SiteEntity struct {
	gorm.Model
	TenentID  string
	Title     string
	UsuarioID string
}
