package handler

import (
	"net/http"

	"github.com/coderj001/phoneguardian/app/model"
	"github.com/jinzhu/gorm"
)

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
	respondJSON(w, http.StatusOK, contacts)
}
