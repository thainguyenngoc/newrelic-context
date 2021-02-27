package main

import (
	"github.com/best-expendables/newrelic-context/nrgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() *gorm.DB {
	dsn := "host=%s user=%s port=%d dbname=%s sslmode=disable password=%s"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	nrgorm.AddGormCallbacks(db)
	return db
}
