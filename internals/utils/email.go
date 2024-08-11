package utils

import (
	"fmt"
	"github.com/Skythrill256/auth-service/internals/config"
	"net/smtp"
)

func SendVerificationEmail(to string, token string, cfg *config.Config) error {
	from := cfg.EmailSender
	password := cfg.EmailPass
	smtpHost := cfg.EmailHost
	smtpPort := cfg.EmailPort

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: Email Verification\n\nPlease verify your email by clicking on the link: %s/verify?token=%s", from, to, "http://localhost:"+cfg.AppPort, token)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	return err
}
