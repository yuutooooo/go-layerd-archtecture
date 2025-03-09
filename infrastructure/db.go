package infrastructure

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
		println("dbUser", dbUser)
		println("dbPass", dbPass)
		println("dbHost", dbHost)
		println("dbPort", dbPort)
		println("dbName", dbName)
		return nil, errors.New("環境変数が設定されていません")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
