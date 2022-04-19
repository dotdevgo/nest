package user

import (
	"fmt"

	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/goava/di"
	"github.com/matcornic/hermes/v2"
)

// UserMailer godoc
type UserMailer struct {
	di.Inject
	nest.Config
}

// SignUp godoc
func (m UserMailer) SignUp(u *User) hermes.Email {
	link := fmt.Sprintf("%s/auth/confirm/%s", m.Config.HTTP.Hostname, u.GetAttribute(AttributeConfirmToken))

	return hermes.Email{
		Body: hermes.Body{
			Name: u.Username,
			Intros: []string{
				"Welcome to MotusApp! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with MotusApp, please click here:",
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
