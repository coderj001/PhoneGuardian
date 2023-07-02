package handler

import (
	"encoding/json"
	"net/http"

	"github.com/coderj001/phoneguardian/app/model"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}

	if user.Name == "" || user.Phone == "" || user.Password == "" {
		respondError(w, http.StatusBadRequest, "empty")
	}

	newUser := model.User{
		Name:        user.Name,
		PhoneNumber: user.Phone,
		Email:       user.Email,
		Password:    user.Password,
	}

	err = db.Create(&newUser).Error
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"user_id": newUser.ID,
	}

	respondJSON(w, http.StatusCreated, response)

}
