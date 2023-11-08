package model

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// modelsToMigrate is a list of models to migrate
var modelsToMigrate = []interface{}{
	&User{},
	&Business{},
	&Category{},
	&BusinessCategory{},
	&Contact{},
	&BusinessLocation{},
}

// InitDB initializes the database connection
func InitDB(username, password, host, port, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db = gormDB
	syncDatabase()
	return db, nil
}

// syncDatabase migrates models to the database
func syncDatabase() {
	err := db.AutoMigrate(modelsToMigrate...)
	if err != nil {
		log.Println(err)
	}
}
