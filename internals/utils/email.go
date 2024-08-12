
package utils

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/Skythrill256/auth-service/internals/config"
)

func SendVerificationEmail(to string, token string, cfg *config.Config) error {
	from := cfg.EmailSender
	appPassword := cfg.SMTP_API_KEY

	subject := "Email Verification"
	body := fmt.Sprintf("Please verify your email by clicking on the link: %s/verify?token=%s", "http://localhost:"+cfg.AppPort, token)

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, appPassword, "smtp.gmail.com")

	// Compose the email.
	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body,
	)

	// Send the email.
	err := smtp.SendMail(
		"smtp.gmail.com:587", // SMTP server address and port.
		auth,
		from,
		[]string{to},
		message,
	)

	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Printf("Email sent from: %s to: %s", from, to)
	return nil
}
