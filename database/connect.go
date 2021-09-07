package database

import (
	"estpguiapi/config"
	"estpguiapi/lib"
	"estpguiapi/models"
	"fmt"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"strconv"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	DB, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn, DefaultStringSize: 256, DisableDatetimePrecision: true, DontSupportRenameIndex: true, DontSupportRenameColumn: true, SkipInitializeWithVersion: false}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
func InitTables() {
	migrateTables := config.Config("MIGRATE")
	db := DB
	if migrateTables == "true" {
		if db.Migrator().HasTable(&models.User{}) == false {
			db.Migrator().CreateTable(&models.User{})
			password, _ := lib.HashPassword("magicword")

			var Users = []models.User{{
				Username:  "admin",
				Email:     "tarik@fink-telecom.com",
				Password:  password,
				FirstName: "Administrator",
				LastName:  "",
				Model:     gorm.Model{},
			}}
			db.Create(&Users)
		}
		if db.Migrator().HasTable(models.Partners{}) == false {
			db.Migrator().CreateTable(&models.Partners{})

		}
		if db.Migrator().HasTable(models.Datacenter{}) == false {
			db.Migrator().CreateTable(&models.Datacenter{})

		}
		if db.Migrator().HasTable(models.EstpInstances{}) == false {
			db.Migrator().CreateTable(&models.EstpInstances{})

		}
		if db.Migrator().HasTable(models.EstpCategories{}) == false {
			db.Migrator().CreateTable(&models.EstpCategories{})

		}
		if db.Migrator().HasTable(models.ESTPIpModel{}) == false {
			db.Migrator().CreateTable(&models.ESTPIpModel{})

		}
		if db.Migrator().HasTable(models.SCTPModel{}) == false {
			db.Migrator().CreateTable(&models.SCTPModel{})

		}
	}
}
