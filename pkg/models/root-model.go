package models

import (
	"GinRestApi/pkg/db"
	"gorm.io/gorm"
)

var ctx *gorm.DB

func init() {
	db.Init()
	ctx = db.Context()
	err := ctx.AutoMigrate(
		&Author{},
		&Book{},
	)

	if err != nil {
		panic("failed to migrate database")
	}
}
