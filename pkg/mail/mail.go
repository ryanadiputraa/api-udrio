package mail

import (
	"fmt"
	"net/smtp"

	"github.com/ryanadiputraa/api-udrio/config"
)

type MailPayload struct {
	Subject string
	Body    string
	To      []string
}

func SendMail(conf config.Mail, payload MailPayload) (err error) {
	auth := smtp.PlainAuth("", conf.Sender, conf.Pass, "smtp.gmail.com")

	msg := fmt.Sprintf("Subject: %s\n%s", payload.Subject, payload.Body)

	err = smtp.SendMail("smtp.gmail.com:587", auth, conf.Sender, payload.To, []byte(msg))
	return
}
