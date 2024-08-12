package handlers

import (
	"encoding/json"
	"github.com/Skythrill256/auth-service/internals/config"
	"github.com/Skythrill256/auth-service/internals/db"
	"github.com/Skythrill256/auth-service/internals/services"
	"github.com/Skythrill256/auth-service/internals/utils"
	"net/http"
)

type Handler struct {
	Repository *db.Repository
	Config     *config.Config
}

func NewHandler(repository *db.Repository, config *config.Config) *Handler {
	return &Handler{
		Repository: repository,
		Config:     config,
	}
}

func (h *Handler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var user utils.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = services.SignUpUser(user, h.Repository, h.Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Registered, Please verify your email"})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var userDTO utils.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	token, err := services.LoginUser(userDTO, h.Repository, h.Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token:": token})

}

func (h *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	err := services.VerifyEmail(token, h.Repository, h.Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email Verified"})
}

func (h *Handler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}
	token, err := services.GoogleLogin(h.Config, h.Repository, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
