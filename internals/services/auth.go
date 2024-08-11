package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/Skythrill256/auth-service/internals/config"
	"github.com/Skythrill256/auth-service/internals/db"
	"github.com/Skythrill256/auth-service/internals/models"
	"github.com/Skythrill256/auth-service/internals/utils"
)

func SignUpUser(user utils.UserDTO, repository *db.Repository, cfg *config.Config) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := &models.User{
		Email:      user.Email,
		Password:   string(hashedPassword),
		IsVerified: false,
	}
	err = repository.CreateUser(newUser)
	if err != nil {
		return err
	}
	verificationToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return err
	}
	return utils.SendVerificationEmail(user.Email, verificationToken, cfg)
}

func LoginUser(user utils.UserDTO, repository *db.Repository, cfg *config.Config) (string, error) {
	storedUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if storedUser == nil {
		return "", errors.New("user not found")
	}

	if !storedUser.IsVerified {
		return "", errors.New("email not verified")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyEmail(token string, repository *db.Repository, cfg *config.Config) error {
	email, err := utils.ParseJWT(token, cfg.JWTSecret)
	if err != nil {
		return err
	}

	return repository.VerifyUserEmail(email)
}
