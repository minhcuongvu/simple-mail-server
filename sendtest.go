package main

import (
	"fmt"
	"os"
)

type MailServer interface {
	SendTestEmail(email *EmailObject) error
}

type MailServerConfig struct {
	MailServerProvider MailServer
}

type MailUser struct {
	email string
	name  string
}

type EmailType int

const (
	PlainText EmailType = iota
	HTML
)

type EmailObject struct {
	subject   string
	body      string
	recipient string
	emailType EmailType
}

func main() {
	// Setup default user
	defaultUser := &MailUser{
		email: fmt.Sprintf(
			"%s@%s",
			"info",
			os.Getenv("EMAIL_FROM_DOMAIN"),
		),
		name: "Michael Vu",
	}

	// Initialize config with SES provider
	config := &MailServerConfig{
		MailServerProvider: &AmazonSESMailServer{
			mailUser:     defaultUser,
			smtpHost:     os.Getenv("SES_SMTP_SERVER"),
			smtpPort:     os.Getenv("SES_SMTP_SERVER_PORT"),
			smtpUsername: os.Getenv("SES_SMTP_USERNAME"),
			smtpPassword: os.Getenv("SES_SMTP_PASSWORD"),
		},
	}

	email := EmailObject{
		subject:   "Test email from AmazonSES",
		body:      "Hello,<br/>This is a test email!",
		recipient: "noedigsti@gmail.com",
		emailType: HTML,
	}

	switch provider := config.MailServerProvider.(type) {
	case *AmazonSESMailServer:
		err := provider.SendTestEmail(&email)
		if err != nil {
			fmt.Printf("Failed to send email: %v\n", err)
		} else {
			fmt.Println("Done.")
		}
	default:
		fmt.Println("Unsupported mail server provider")
	}
}
