package utils

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/Skythrill256/auth-service/internals/config"
)

func SendVerificationEmail(to string, token string, cfg *config.Config) error {
	from := cfg.EmailSender
	appPassword := cfg.EmailPass

	subject := "Email Verification"
	body := fmt.Sprintf(`
		<html>
		<body>
			<p>Please verify your email by clicking on the link:</p>
			<a href="http://localhost:%s/verify-email?token=%s">Verify Email</a>
		</body>
		</html>
	`, cfg.AppPort, token)

	auth := smtp.PlainAuth("", from, appPassword, "smtp.gmail.com")

	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body,
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
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
