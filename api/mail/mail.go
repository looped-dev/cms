package mail

import (
	"fmt"
	"os"
	"strconv"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendEmail(emailTo, emailFrom, subject, body string) error {
	// todo: use viper to get these values
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST")

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		return fmt.Errorf("Error parsing MAIL_PORT: %v", err)
	}

	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = port
	server.Username = username
	server.Password = password
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From Me <me@host.com>")
	email.AddTo("you@example.com")
	email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, body)
	email.AddAttachment("super_cool_file.png")

	return email.Send(smtpClient)
}
