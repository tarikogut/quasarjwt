package models

import (
	"encoding/gob"
	"gorm.io/gorm"
	"time"
)

type ESTPIpModel struct {
	Ipaddress string `json:"ipaddress"`
	AddressType bool `json:"address_type"`
	ESTPId int64 `json:"estp_id"`

	gorm.Model
}

func (ESTPIpModel *ESTPIpModel) TableName() string {
	return "estpgui_sctp_ipaddress"
}

func (ESTPIpModel *ESTPIpModel) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("created_at", time.Now())
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}

func (ESTPIpModel *ESTPIpModel) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}
func init() {
	gob.Register(ESTPIpModel{})
}
