package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailer struct {
	fromEmail string
	apikey    string
	client    *sendgrid.Client
}

func NewSendGrid(fromEmail, apikey string) *SendGridMailer {
	return &SendGridMailer{
		fromEmail: fromEmail,
		apikey:    apikey,
		client:    sendgrid.NewSendClient(apikey),
	}
}

func (m *SendGridMailer) Send(templateFile, username, email string, data any, isSandbox bool) error {
	from := mail.NewEmail(FromName, m.fromEmail)
	to := mail.NewEmail(username, email)

	// template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return err
	}

	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())

	message.SetMailSettings(&mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &isSandbox,
		},
	})

	for i := 0; i < maxRetries; i++ {
		response, err := m.client.Send(message)
		if err != nil {
			log.Printf("Error sending email: %v, attemp %d of %d", email, i+1, maxRetries)
			log.Printf("Error: %v", err.Error())

			// exponential backoff
			time.Sleep(time.Second * time.Duration(i*i))
			continue
		}

		log.Printf("Email sent to %s, response: %v", email, response.StatusCode)
		return nil
	}

	return fmt.Errorf("failed to send email after %d attemps", maxRetries)
}
