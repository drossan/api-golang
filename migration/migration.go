package migration

import (
	"../configuration"
	"../models"
)

// Migrate - created tables in db
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Level{})
	db.Model(&models.User{}).Related(&models.Level{}, "Level")
}
