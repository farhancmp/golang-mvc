package models

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DbConnect() (db *gorm.DB, sql *sql.DB, err error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv().DBUsername,
		GetEnv().DBPassword,
		GetEnv().DBHost,
		GetEnv().DBPort,
		GetEnv().DBName)

	db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Println(err.Error())
	}
	sql, err = db.DB()
	if err != nil {
		log.Println(err.Error())
	}

	return db, sql, err
}
