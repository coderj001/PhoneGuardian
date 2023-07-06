package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coderj001/phoneguardian/app/auth"
	"github.com/coderj001/phoneguardian/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	// "golang.org/x/crypto/bcrypt"
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
		RespondError(w, http.StatusBadRequest, err.Error())
	}

	if user.Name == "" || user.Phone == "" || user.Password == "" {
		RespondError(w, http.StatusBadRequest, "empty")
	}

	newUser := model.User{
		Name:        user.Name,
		PhoneNumber: user.Phone,
		Email:       user.Email,
		Password:    user.Password,
	}

	err = db.Create(&newUser).Error
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"user_id": newUser.ID,
	}

	respondJSON(w, http.StatusCreated, response)

}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

func LoginUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	user := model.User{}
	err = db.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// TODO: will added logic later
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if user.Password != request.Password {
		RespondError(w, http.StatusUnauthorized, "Invalid phone or password")
		return
	}

	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID)

	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := LoginResponse{
		Token:  token,
		UserID: user.ID,
	}

	respondJSON(w, http.StatusOK, response)

}

type UserDetailResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func GetUserDetailes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	fmt.Println(userID)

	var user model.User
	if err := db.Find(&user, userID).Error; err != nil {
		RespondError(w, http.StatusNotFound, "User not found")
		return
	}

	response := UserDetailResponse{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}
	respondJSON(w, http.StatusOK, response)

}
