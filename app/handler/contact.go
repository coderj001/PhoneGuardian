package handler

import (
	"encoding/json"
	"net/http"

	"github.com/coderj001/phoneguardian/app/model"
	"github.com/jinzhu/gorm"
)

type CreateContactRequest struct {
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type CreateContactResponse struct {
	ContactID uint `json:"contact_id"`
}

func CreateContact(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var request CreateContactRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	if request.UserID <= 0 || request.Name == "" || request.PhoneNumber == "" {
		respondError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	contact := model.Contact{
		UserID:      request.UserID,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
	}

	err = db.Create(&contact).Error

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create contact")
		return
	}

	response := CreateContactResponse{
		ContactID: contact.ID,
	}

	respondJSON(w, http.StatusCreated, response)
}
