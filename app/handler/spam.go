package handler

import (
	"encoding/json"
	"net/http"

	"github.com/coderj001/phoneguardian/app/model"
	"github.com/jinzhu/gorm"
)

type SpamRequest struct {
	PhoneNumber string `json:"phone_number"`
}

func MarkNumberAsSpam(db *gorm.DB, w http.ResponseWriter, r *http.Request)  {
	// userID := r.Context().Value("userID").(uint)

	var requestBody SpamRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	defer r.Body.Close()
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	phoneNumber := requestBody.PhoneNumber

	spam := model.Spam{
		PhoneNumber: phoneNumber,
	}
	err = db.Create(&spam).Error
	
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Failed to mark number as spam")
		return
	}

	// Respond with success status
	respondJSON(w, http.StatusCreated, nil)
}
