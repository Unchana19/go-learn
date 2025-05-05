package mailer

import "embed"

const (
	FromName   = "Social Media"
	maxRetries = 3
	// Templates
	UserWelcomeTemplate = "user_invitation.templ"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}
