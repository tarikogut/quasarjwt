package models

import (
	"encoding/gob"
	"gorm.io/gorm"
	"time"
)

type EstpInstances struct {
	OrganizationsId int64  `json:"organizations_id"`
	DatacenterId    int64  `json:"datacenter_id"`
	Hostname        string `json:"hostname"`
	Port            int64  `json:"port"`
	ApiUsername     string `json:"api_username"`
	ApiPassword     string `json:"api_password"`
	Https           bool   `json:"https"`
	gorm.Model
}
func (EstpInstances *EstpInstances) TableName() string {
	return "estpgui_estp_servers"
}

func (EstpInstances *EstpInstances) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("created_at", time.Now())
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}

func (EstpInstances *EstpInstances) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}
func init() {
	gob.Register(EstpInstances{})
}