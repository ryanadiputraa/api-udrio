package mail

import (
	"fmt"
	"net/smtp"

	"github.com/spf13/viper"
)

func SendMail(subject string, body string, to []string) (err error) {
	auth := smtp.PlainAuth("", viper.GetString("MAIL_SENDER"), viper.GetString("MAIL_SENDER_PASS"), "smtp.gmail.com")

	msg := fmt.Sprintf("Subject: %s\n%s", subject, body)

	err = smtp.SendMail("smtp.gmail.com:587", auth, viper.GetString("MAIL_SENDER"), to, []byte(msg))
	return
}
