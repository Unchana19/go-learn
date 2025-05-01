package mailer

import (
	"bytes"
	"errors"
	"html/template"
)

type mailtrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailtrapClient(fromEmail, apiKey string) (*mailtrapClient, error) {
	if apiKey == "" {
		return nil, errors.New("mailtrap api key is required")
	}

	return &mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}

func (m *mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err
	}

	return 200, nil
}
