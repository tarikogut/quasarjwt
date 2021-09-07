package models

import (
	"encoding/gob"
	"gorm.io/gorm"
	"time"
)

type SCTPModel struct {
	Name       string `json:"name"`
	Enable     bool   `json:"enable"`
	LogLevel   int    `json:"log-level"`
	LogFile    string `json:"log-file"`
	LocalIp    string `json:"local-ip"`
	RemoteIp   string `json:"remote-ip"`
	LocalPort  int    `json:"local-port"`
	RemotePort int    `json:"remote-port"`
	Passive    bool   `json:"passive"`
	Heartbeat  int    `json:"heartbeat"`
	EstpId     int64  `json:"estp_id"`
	Mtu        int    `json:"mtu"`
	gorm.Model
}

func (SCTPModel *SCTPModel) TableName() string {
	return "estpgui_sctpconnections"
}

func (SCTPModel *SCTPModel) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("created_at", time.Now())
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}

func (SCTPModel *SCTPModel) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}
func init() {
	gob.Register(SCTPModel{})
}
