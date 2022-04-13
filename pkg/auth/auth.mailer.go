package auth

import (
	"fmt"

	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/dotdevgo/nest/pkg/user"

	"github.com/goava/di"
	"github.com/matcornic/hermes/v2"
)

type AuthMailer struct {
	di.Inject
	nest.Config
}

// Restore godoc
func (m *AuthMailer) Restore(u *user.User) hermes.Email {

	link := fmt.Sprintf(
		"%s/auth/reset/%s/%s",
		m.Config.HTTP.Hostname,
		u.UUID,
		u.GetAttribute(user.AttributeResetToken),
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
func (m *AuthMailer) ResetToken(event user.EventResetToken) hermes.Email {
	link := m.Config.HTTP.Hostname

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
