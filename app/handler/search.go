package handler

import (
	"net/http"

	"github.com/coderj001/phoneguardian/app/model"
	"github.com/jinzhu/gorm"
)

type ContactResult struct {
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	SpamLikelihood string `json:"spam_likelihood"`
}

func SeachContact(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	phone_number := r.URL.Query().Get("phone_number")

	if name == "" && phone_number == "" {
		RespondError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}

	var contacts []model.Contact

	err := db.Where("name LIKE ? OR phone_number = ?", "%"+name+"%", phone_number).Find(&contacts).Error
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Failed to search contacts")
		return
	}
	var results []ContactResult
	for _, contact := range contacts {
		result := ContactResult{
			Name:           contact.Name,
			PhoneNumber:    contact.PhoneNumber,
			SpamLikelihood: calculateSpamLikeliHood(db, contact.PhoneNumber),
		}
		results = append(results, result)
	}
	respondJSON(w, http.StatusOK, results)
}

func calculateSpamLikeliHood(db *gorm.DB, phoneNumber string) string {
	var count int
	db.Model(&model.Spam{}).Where("phone_number = ?", phoneNumber).Count(&count)

	switch {
	case count < 5:
		return "Low"
	case count >= 5 && count < 10:
		return "Medium"
	default:
		return "High"
	}
}
