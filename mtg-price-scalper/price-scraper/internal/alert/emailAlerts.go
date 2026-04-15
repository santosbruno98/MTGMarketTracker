package alert

import (
	"context"

	gomail "gopkg.in/gomail.v2"
)

type EmailAlertServiceImpl struct {
	Server  string
	Port    int
	Sender  string
	To      string
	Subject string
}

const (
	SERVER       = "smtp.gmail.com"
	PORT         = 587
	FROM_EMAIL   = "bs745645@gmail.com"
	APP_PASSWORD = "toCreate" // TODO: go to gmail and create an app password for this
	SUBJECT      = "Alert from mtg-price-scalper"
)

func NewEmailAlertServiceImpl(toEmail string) *EmailAlertServiceImpl {
	return &EmailAlertServiceImpl{
		Server:  SERVER,
		Port:    PORT,
		Sender:  FROM_EMAIL,
		To:      toEmail,
		Subject: SUBJECT,
	}
}

func (emailService *EmailAlertServiceImpl) SendEmailAlert(ctx context.Context, toEmail string, message string) error {
	// make the email dial

	newMail := gomail.NewMessage()

	newMail.SetHeader("From", emailService.Sender)
	newMail.SetHeader("To", emailService.To)
	newMail.SetHeader("Subject", emailService.Subject)
	newMail.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		emailService.Server,
		emailService.Port,
		emailService.Sender,
		APP_PASSWORD)
	return dialer.DialAndSend(newMail)
}
