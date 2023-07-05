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
	userID := r.Context().Value("userID").(uint)

	var request CreateContactRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	if request.Name == "" || request.PhoneNumber == "" {
		RespondError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	contact := model.Contact{
		UserID:      userID,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
	}

	err = db.Create(&contact).Error

	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Failed to create contact")
		return
	}

	response := CreateContactResponse{
		ContactID: contact.ID,
	}

	respondJSON(w, http.StatusCreated, response)
}
