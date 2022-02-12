package emails

import (
	"fmt"
	"strconv"

	"github.com/looped-dev/cms/api/utils/configs"
	mail "github.com/xhit/go-simple-mail/v2"
)

type SendMailConfig struct {
	emailTo   string
	emailFrom string
	subject   string
	htmlBody  string
	plainBody string
}

func NewSMTPClient() (*mail.SMTPClient, error) {
	// todo: use viper to get these values
	username := configs.GetConfig("MAIL_USERNAME")
	password := configs.GetConfig("MAIL_PASSWORD")
	host := configs.GetConfig("MAIL_HOST")
	port, err := strconv.Atoi(configs.GetConfig("MAIL_PORT"))
	if err != nil {
		return nil, fmt.Errorf("Error parsing MAIL_PORT: %v", err)
	}
	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = port
	server.Username = username
	server.Password = password
	server.Encryption = mail.EncryptionTLS
	return server.Connect()
}

func SendEmail(smtpClient *mail.SMTPClient, config SendMailConfig) error {
	// Create email
	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("From Me <%s>", config.emailFrom))
	email.AddTo(config.emailTo)
	email.SetSubject(config.subject)
	email.SetBody(mail.TextHTML, config.htmlBody)
	email.SetBody(mail.TextPlain, config.plainBody)
	return email.Send(smtpClient)
}
