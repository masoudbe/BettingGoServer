package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type VoteInfo struct {
	gorm.Model
	Name string
	Vote string
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&VoteInfo{})
	return db
}
