package models

import (
	"encoding/gob"
	"gorm.io/gorm"
	"time"
)

type Partners struct {
	PartnersName string `json:"partners_name"`
	gorm.Model
}

func (Partners *Partners) TableName() string {
	return "estpgui_estp_partners"
}

func (Partners *Partners) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("created_at", time.Now())
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}

func (Partners *Partners) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}
func init() {
	gob.Register(Partners{})
}
