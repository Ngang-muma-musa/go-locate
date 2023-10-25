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
}

// InitDB initializes the database connection
func InitDB(username, password, host, port, dbName string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	log.Println(dsn)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = gormDB
	syncDatabase()
	return nil
}

func ClearTables(models ...any) {
	for _, model := range models {
		switch model.(type) {
		case *User:
			db.Exec("truncate table `users`")
		case *Business:
			db.Exec("truncate table `businesses`")
		case *BusinessCategory:
			db.Exec("truncate table `busineess_category`")
		case *Category:
			db.Exec("truncate table `categories`")
		case *Contact:
			db.Exec("truncate table `contacts`")
		}
	}
}

// CloseDB closes the underlying database connection.
func CloseDB() {
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

// syncDatabase migrates models to the database
func syncDatabase() {
	err := db.AutoMigrate(modelsToMigrate...)
	if err != nil {
		log.Println(err)
	}
}
