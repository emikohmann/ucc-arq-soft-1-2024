package clients

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/dao/users"
)

const (
	connectionString = "%s:%s@tcp(%s:%d)/:%s?charset=utf8&parseTime=True"
)

var (
	DBClient *gorm.DB
)

func init() {
	dbName := "final"
	dbUser := "root"
	dbPassword := "1234"
	dbHost := "localhost"
	dbPort := 3306

	connection := fmt.Sprintf(connectionString, dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to DB: %v", err))
	}
	DBClient = db
}

func StartDB() {
	var user users.User
	if err := DBClient.AutoMigrate(&user); err != nil {
		panic(fmt.Errorf("error creating entities: %v", err))
	}
}
