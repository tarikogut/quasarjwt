package models

import (
	"encoding/gob"
	"gorm.io/gorm"
	"time"
)

type Datacenter struct {
	OrganizationId int64  `json:"organization_id"`
	DatacenterName string `json:"datacenter_name"`
	gorm.Model
}

func (Datacenter *Datacenter) TableName() string {
	return "estpgui_datacenters"
}

func (Datacenter *Datacenter) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("created_at", time.Now())
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}

func (Datacenter *Datacenter) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}
func init() {
	gob.Register(Datacenter{})
}
