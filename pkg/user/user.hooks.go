package user

import (
	"context"
	"log"

	"github.com/dotdevgo/nest/pkg/mailer"
	"github.com/goava/di"
	"github.com/mustafaturan/bus/v3"
)

// UserHooks godoc
type UserHooks struct {
	di.Inject

	*UserMailer
	*mailer.Mailer
}

// EventUserSignUp godoc
func (h UserHooks) EventUserSignUp() bus.Handler {
	return bus.Handler{
		Matcher: EventUserSignUp,
		Handle: func(ctx context.Context, e bus.Event) {
			u, ok := e.Data.(*User)
			if !ok {
				return
			}

			go func() {
				template := h.UserMailer.SignUp(u)
				m, err := h.Mailer.NewEmail(template)
				if err != nil {
					log.Print(err)
					return
				}

				m.To = []string{u.Email}
				m.Subject = "NestApp: Confirm your account"

				if err := h.Mailer.Send(m); err != nil {
					log.Print(err)
					return
				}
			}()
		},
	}
}

// from := "bot@dotdev.ru"

// user := "bot@dotdev.ru"
// password := "dotdevWbz1=5"

// to := []string{
// 	"me@gotgame.ru",
// }

// addr := "smtp.yandex.ru:25"
// host := "smtp.yandex.ru"

// msg := []byte("From: bot@dotdev.ru\r\n" +
// 	"To: me@gotgame.ru\r\n" +
// 	"Subject: Test mail\r\n\r\n" +
// 	"Email body\r\n")

// auth := smtp.PlainAuth("", user, password, host)

// err := smtp.SendMail(addr, auth, from, to, msg)

// if err != nil {
// 	log.Print(err)
// 	return
// }

// log.Print("Email sent successfully")

// addr := "smtp.yandex.ru:25"
// host := "smtp.yandex.ru"

// auth := smtp.PlainAuth("", "bot@dotdev.ru", "dotdevWbz1=5", host)

// if err := m.Send(addr, auth); err != nil {
// 	log.Print(err)
// }
