package emails

import (
	"fmt"
	"strconv"

	"github.com/looped-dev/cms/api/utils/configs"
	mail "github.com/xhit/go-simple-mail/v2"
)

type SendMailConfig struct {
	EmailTo   string
	EmailFrom string
	Subject   string
	HtmlBody  string
	PlainBody string
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

func NewMockSMTPClient(host string, port int) (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = port
	return server.Connect()
}

func SendEmail(smtpClient *mail.SMTPClient, config SendMailConfig) error {
	// Create email
	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("From Me <%s>", config.EmailFrom))
	email.AddTo(config.EmailTo)
	email.SetSubject(config.Subject)
	email.SetBody(mail.TextHTML, config.HtmlBody)
	email.SetBody(mail.TextPlain, config.PlainBody)
	return email.Send(smtpClient)
}
