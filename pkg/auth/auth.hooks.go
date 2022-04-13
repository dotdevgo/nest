package auth

import (
	"context"
	"log"

	"github.com/dotdevgo/nest/pkg/mailer"
	"github.com/dotdevgo/nest/pkg/user"

	"github.com/goava/di"
	"github.com/mustafaturan/bus/v3"
)

// AuthHooks godoc
type AuthHooks struct {
	di.Inject

	*AuthMailer
	*mailer.Mailer
}

// EventRestore godoc
func (h *AuthHooks) EventRestore() bus.Handler {
	return bus.Handler{
		Matcher: user.EventUserRestore,
		Handle: func(ctx context.Context, e bus.Event) {
			u, ok := e.Data.(*user.User)
			if !ok {
				return
			}

			log.Printf("AuthHooks: OnRestore %v", u.ID)

			go func() {
				template := h.AuthMailer.Restore(u)
				m, err := h.Mailer.NewEmail(template)
				if err != nil {
					log.Print(err)
					return
				}

				m.To = []string{u.Email}
				m.Subject = "Reset account password"

				if err := h.Mailer.Send(m); err != nil {
					log.Print(err)
					return
				}
			}()
		},
	}
}

// EventResetToken godoc
func (h *AuthHooks) EventResetToken() bus.Handler {
	return bus.Handler{
		Matcher: user.EventUserResetToken,
		Handle: func(ctx context.Context, e bus.Event) {
			event, ok := e.Data.(user.EventResetToken)
			if !ok {
				return
			}

			u := event.User
			log.Printf("AuthHooks: OnResetToken %v", u.ID)

			go func() {
				template := h.AuthMailer.ResetToken(event)
				m, err := h.Mailer.NewEmail(template)
				if err != nil {
					log.Print(err)
					return
				}

				m.To = []string{u.Email}
				m.Subject = "Password has been reset"

				if err := h.Mailer.Send(m); err != nil {
					log.Print(err)
					return
				}
			}()
		},
	}
}

// if "" == os.Getenv("SMTP_HOST") {
// 	log.Printf("AuthHooks: Skip SMTP invalid %v", u.ID)
// 	return
// }

// go func() {
// 	email := h.Hermes.Restore(u)

// 	m, err := mail.Prepare(h.Hermes.Hermes, email)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	m.To = []string{u.Email}
// 	m.Subject = "Reset account password"

// 	if err := mail.Send(m); err != nil {
// 		log.Print(err)
// 		return
// 	}
// }()
// if "" == os.Getenv("SMTP_HOST") {
// 	log.Printf("AuthHooks: Skip SMTP invalid %v", u.ID)
// 	return
// }

// go func() {
// 	email := h.Hermes.ResetToken(event)

// 	m, err := mail.Prepare(h.Hermes.Hermes, email)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	m.To = []string{u.Email}
// 	m.Subject = "Password has been reset"

// 	if err := mail.Send(m); err != nil {
// 		log.Print(err)
// 		return
// 	}
// }()
