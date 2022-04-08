package database

import (
	"fmt"
	"time"

	"test-majoo-api/config"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBInit(c *config.Config) *gorm.DB {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName)
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(5)

	db.LogMode(true)
	return db
}
