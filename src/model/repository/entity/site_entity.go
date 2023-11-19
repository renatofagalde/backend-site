package entity

import "gorm.io/gorm"

type SiteEntity struct {
	gorm.Model
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
