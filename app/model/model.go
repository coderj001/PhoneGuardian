package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// User represents a user model
type User struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" gorm:"unique_index"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Contacts []*Contact `json:"contacts" gorm:"foreignkey:UserID"`
}

// Contact represents a contact model
type Contact struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" gorm:"unique_index:idx_user_phone_number"`
	UserID      uint   `json:"-"`
}

// Spam represents a spam number model
type Spam struct {
	gorm.Model
	PhoneNumber string `json:"phone_number" gorm:"unique_index"`
}

// SearchResult represents a search result model
type SearchResult struct {
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	SpamLikelihood int    `json:"spam_likelihood"`
	IsRegistered   bool   `json:"is_registered"`
	Email          string `json:"email,omitempty"`
}

// DBMigrate will handle database migration
func DBMigrate(db *gorm.DB) *gorm.DB {
	log.Println("[+] database migration")
	db.AutoMigrate(&User{}, &Contact{}, &Spam{})
	db.Model(&Contact{}).AddUniqueIndex("idx_user_phone_number", "user_id", "phone_number")
	return db
}
