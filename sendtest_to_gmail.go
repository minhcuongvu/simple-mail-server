package main

import (
	"fmt"
	"net/smtp"
)

type AmazonSES interface {
	SendTestEmail(emailObject *EmailObject) error
}

type AmazonSESMailServer struct {
	mailUser     *MailUser
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
}

// SendTestEmail method for AmazonSESMailServer with support for PlainText and HTML formats
func (s *AmazonSESMailServer) SendTestEmail(emailObject *EmailObject) error {
	from := fmt.Sprintf("\"%s\" <%s>", s.mailUser.name, s.mailUser.email)

	var msg string
	switch emailObject.emailType {
	case PlainText:
		// Plain Text Email
		msg = fmt.Sprintf(
			"From: %s\nTo: %s\nSubject: %s\n\n%s",
			from,
			emailObject.recipient,
			emailObject.subject,
			emailObject.body,
		)
	case HTML:
		// HTML Email
		msg = fmt.Sprintf(
			"From: %s\nTo: %s\nSubject: %s\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s",
			from, emailObject.recipient, emailObject.subject, emailObject.body,
		)
	default:
		return fmt.Errorf("unsupported email type")
	}

	fmt.Println(msg)

	// Authenticate using SES credentials
	auth := smtp.PlainAuth("", s.smtpUsername, s.smtpPassword, s.smtpHost)

	// Send the email via SES SMTP server
	err := smtp.SendMail(
		s.smtpHost+":"+s.smtpPort,
		auth,
		s.mailUser.email,
		[]string{emailObject.recipient},
		[]byte(msg),
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
