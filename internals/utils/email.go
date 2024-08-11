package utils

import (
	"fmt"
	"github.com/Skythrill256/auth-service/internals/config"
	"gopkg.in/gomail.v2"
	"log"
)

func SendVerificationEmail(to string, token string, cfg *config.Config) error {
	from := cfg.EmailSender
	apiKey := cfg.SMTP_KEY

	subject := "Email Verification"
	body := fmt.Sprintf("Please verify your email by clicking on the link: %s/verify?token=%s", "http://localhost:"+cfg.AppPort, token)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp-relay.sendinblue.com", 587, from, apiKey)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
}
