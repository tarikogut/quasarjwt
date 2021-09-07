package models

import (
	"encoding/gob"
	"gorm.io/gorm"
	"time"
)

type EstpCategories struct {
	CategoryName string `json:"category_name"`
	gorm.Model
}
func (EstpCategories *EstpCategories) TableName() string {
	return "estpgui_connection_categories"
}

func (EstpCategories *EstpCategories) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("created_at", time.Now())
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}

func (EstpCategories *EstpCategories) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}
func init() {
	gob.Register(EstpCategories{})
}
