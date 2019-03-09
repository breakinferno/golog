package db

import (
	"github.com/breakinferno/golog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() (*gorm.DB, error) {
	var err error
	db, err = Connect()
	return db, err
}

func ConnectDB(option string) (*gorm.DB, error) {
	// db, err := gorm.Open(“mysql”, “user:password@/dbname?charset=utf8&parseTime=True&loc=Local”)
	db, err := gorm.Open("mysql", option)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Connect() (*gorm.DB, error) {
	user := config.GetString("db.options.user")
	password := config.GetString("db.options.password")
	host := config.GetString("db.options.host")
	port := config.GetString("db.options.port")
	dbname := config.GetString("db.options.dbname")
	connectTimeout := config.GetString("db.options.connectTimeout")

	db, err := ConnectDB(user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + dbname + "?charset=utf8&parseTime=True&loc=Local&timeout=" + connectTimeout)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}
