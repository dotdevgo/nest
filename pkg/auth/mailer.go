package auth

import (
	"fmt"

	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/user"

	"github.com/goava/di"
	"github.com/matcornic/hermes/v2"
)

type AuthMailer struct {
	di.Inject
	nest.Config
}

// SignUp godoc
func (m AuthMailer) SignUp(u user.User) hermes.Email {
	link := fmt.Sprintf(
		"%s/auth/confirm/%s",
		m.Config.HTTP.Hostname,
		u.GetAttribute(AttributeConfirmToken),
	)

	return hermes.Email{
		Body: hermes.Body{
			Name: u.Username,
			Intros: []string{
				fmt.Sprintf("Welcome to %s! We're very excited to have you on board.", m.Config.App.Name),
			},
			Actions: []hermes.Action{
				{
					Instructions: fmt.Sprintf("To get started with %s, please click here:", m.Config.App.Name),
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  link,
					},
				},
			},
			// Outros: []string{
			// 	"Need help, or have questions? Just reply to this email, we'd love to help.",
			// },
		},
	}
}

// Restore godoc
func (m AuthMailer) Restore(u user.User) hermes.Email {
	link := fmt.Sprintf(
		"%s/auth/reset/%s/%s",
		m.Config.HTTP.Hostname,
		u.ID,
		u.GetAttribute(AttributeResetToken),
	)

	return hermes.Email{
		Body: hermes.Body{
			Name: u.Username,
			Intros: []string{
				"Reset password for your account.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To reset your password, please click here:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Reset password",
						Link:  link,
					},
				},
			},
		},
	}
}

// ResetToken godoc
func (m AuthMailer) ResetToken(event EventResetToken) hermes.Email {
	link := fmt.Sprintf("%s/auth/signin", m.Config.CORS.Origin)

	return hermes.Email{
		Body: hermes.Body{
			Name: event.User.Username,
			Intros: []string{
				fmt.Sprintf("New password: %s", event.Password),
			},
			Actions: []hermes.Action{
				{
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Sign in to account",
						Link:  link,
					},
				},
			},
		},
	}
}

// ResetEmail godoc
func (m AuthMailer) ResetEmail(u user.User) hermes.Email {
	link := fmt.Sprintf(
		"%s/auth/confirm/%s",
		m.Config.HTTP.Hostname,
		u.GetAttribute(AttributeConfirmToken),
	)

	return hermes.Email{
		Body: hermes.Body{
			Name: u.Username,
			Intros: []string{
				"You have changed your email.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To confirm new Email click here:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your email",
						Link:  link,
					},
				},
			},
			// Outros: []string{
			// 	"Need help, or have questions? Just reply to this email, we'd love to help.",
			// },
		},
	}
}
