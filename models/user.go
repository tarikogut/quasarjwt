package models

import (
	"encoding/gob"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Username  string `gorm:"unique_index;not null" json:"username"`
	Email     string `gorm:"unique_index;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"LastName"`
	gorm.Model
}

func (User *User) TableName() string {
	return "estp_users"
}

func (User *User) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("created_at", time.Now())
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}

func (User *User) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("updated_at", time.Now())
	return nil
}
func init() {
	gob.Register(User{})
}
