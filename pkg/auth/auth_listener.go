package auth

import (
	"context"

	// "dotdev/nest/pkg/logger"
	"dotdev/nest/pkg/mailer"

	"github.com/defval/di"
	"github.com/mustafaturan/bus/v3"
)

// AuthListener godoc
type AuthListener struct {
	di.Inject

	*AuthMailer
	*mailer.Mailer
}

func (h AuthListener) SignUp(ctx context.Context, e bus.Event) {
	event, ok := e.Data.(EventAuthGeneric)
	if !ok {
		return
	}

	go func() {
		template := h.AuthMailer.SignUp(*event.User)
		m, err := h.Mailer.NewEmail(template)
		if err != nil {
			return
		}

		m.To = []string{event.User.Email}
		m.Subject = "Confirm your account"

		if err := h.Mailer.Send(m); err != nil {
			return
		}
	}()
}

func (h AuthListener) Restore(ctx context.Context, e bus.Event) {
	event, ok := e.Data.(EventAuthGeneric)
	if !ok {
		return
	}

	// logger.Log("AuthHooks: OnRestore %v", u.ID)

	go func() {
		template := h.AuthMailer.Restore(*event.User)
		m, err := h.Mailer.NewEmail(template)
		if err != nil {
			return
		}

		m.To = []string{event.User.Email}
		m.Subject = "Reset account password"

		if err := h.Mailer.Send(m); err != nil {
			return
		}
	}()
}

func (h AuthListener) ResetToken(ctx context.Context, e bus.Event) {
	event, ok := e.Data.(EventResetToken)
	if !ok {
		return
	}

	go func() {
		template := h.AuthMailer.ResetToken(event)
		m, err := h.Mailer.NewEmail(template)
		if err != nil {
			return
		}

		m.To = []string{event.User.Email}
		m.Subject = "Password has been reset"

		if err := h.Mailer.Send(m); err != nil {
			return
		}
	}()
}

func (h AuthListener) ResetEmail(ctx context.Context, e bus.Event) {
	event, ok := e.Data.(EventAuthGeneric)
	if !ok {
		return
	}

	// logger.Log("AuthHooks: ChangeEmail %v", u.ID)

	go func() {
		template := h.AuthMailer.ResetEmail(*event.User)
		m, err := h.Mailer.NewEmail(template)
		if err != nil {
			return
		}

		m.To = []string{event.User.Email}
		m.Subject = "Confirm your new email"

		if err := h.Mailer.Send(m); err != nil {
			return
		}
	}()
}

// EventRestore godoc
// func (h AuthHooks) EventRestore() bus.Handler {
// 	return bus.Handler{
// 		Matcher: user.EventUserRestore,
// 		Handle: ,
// 	}
// }

// EventResetToken godoc
// func (h AuthHooks) EventResetToken() bus.Handler {
// 	return bus.Handler{
// 		Matcher: user.EventUserResetToken,
// Handle: func(ctx context.Context, e bus.Event) {
// 	event, ok := e.Data.(user.EventResetToken)
// 	if !ok {
// 		return
// 	}

// 	u := event.User
// 	logger.Log("AuthHooks: OnResetToken %v", u.ID)

// 	go func() {
// 		template := h.AuthMailer.ResetToken(event)
// 		m, err := h.Mailer.NewEmail(template)
// 		if err != nil {
// 			logger.Log(err)
// 			return
// 		}

// 		m.To = []string{u.Email}
// 		m.Subject = "Password has been reset"

// 		if err := h.Mailer.Send(m); err != nil {
// 			logger.Log(err)
// 			return
// 		}
// 	}()
// },
// 	}
// }

// if "" == os.Getenv("SMTP_HOST") {
// 	logger.Log("AuthHooks: Skip SMTP invalid %v", u.ID)
// 	return
// }

// go func() {
// 	email := h.Hermes.Restore(u)

// 	m, err := mail.Prepare(h.Hermes.Hermes, email)
// 	if err != nil {
// 		logger.Log(err)
// 		return
// 	}

// 	m.To = []string{u.Email}
// 	m.Subject = "Reset account password"

// 	if err := mail.Send(m); err != nil {
// 		logger.Log(err)
// 		return
// 	}
// }()
// if "" == os.Getenv("SMTP_HOST") {
// 	logger.Log("AuthHooks: Skip SMTP invalid %v", u.ID)
// 	return
// }

// go func() {
// 	email := h.Hermes.ResetToken(event)

// 	m, err := mail.Prepare(h.Hermes.Hermes, email)
// 	if err != nil {
// 		logger.Log(err)
// 		return
// 	}

// 	m.To = []string{u.Email}
// 	m.Subject = "Password has been reset"

// 	if err := mail.Send(m); err != nil {
// 		logger.Log(err)
// 		return
// 	}
// }()
