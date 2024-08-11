package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Skythrill256/auth-service/internals/config"
	"github.com/Skythrill256/auth-service/internals/db"
	"github.com/Skythrill256/auth-service/internals/models"
	"github.com/Skythrill256/auth-service/internals/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.GoogleRedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleLogin(cfg *config.Config, repository *db.Repository, code string) (string, error) {
	oauthConfig := GetGoogleOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var googleUser map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return "", err
	}

	email, ok := googleUser["email"].(string)
	if !ok {
		return "", errors.New("failed to get email from Google response")
	}
	googleID, ok := googleUser["sub"].(string)
	if !ok {
		return "", errors.New("failed to get Google ID from response")
	}

	user, err := repository.GetUserByGoogleID(googleID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      email,
			IsVerified: true,
			GoogleID:   &googleID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
