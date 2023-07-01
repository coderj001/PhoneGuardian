package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// DBMigrate will handle database migration
func DBMigrate(db *gorm.DB) *gorm.DB {
	log.Println("[+] Database migration...")
	db.AutoMigrate()
	return db
}
