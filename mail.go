package main

import (
	"bytes"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func sendErrorMail(err error) {
	var errmsg bytes.Buffer

	errmsg.WriteString("QReports encountered the following error: ")
	errmsg.WriteString(err.Error())

	e := email.NewEmail()
	e.From = settings.SupportEmail
	e.To = []string{settings.ToEmail}
	e.Subject = "QReports had an error"
	e.Text = []byte(errmsg.String())

	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", settings.SupportEmail, settings.EmailPass, "smtp.gmail.com"))
}
