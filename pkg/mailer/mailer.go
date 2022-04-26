package mailer

import (
	"fmt"
	"net/smtp"

	"dotdev/nest/pkg/logger"
	"dotdev/nest/pkg/nest"

	"github.com/goava/di"
	"github.com/jordan-wright/email"
	"github.com/matcornic/hermes/v2"
)

type Mailer struct {
	di.Inject
	nest.Config
	*hermes.Hermes
}

// NewEmail godoc
func (c Mailer) NewEmail(template hermes.Email) (*email.Email, error) {
	m := email.NewEmail()

	body, err := c.Hermes.GenerateHTML(template)
	if err != nil {
		logger.Log(err)
		return nil, err
	}

	text, err := c.Hermes.GeneratePlainText(template)
	if err != nil {
		logger.Log(err)
		return nil, err
	}

	m.Text = []byte(text)
	m.HTML = []byte(body)

	return m, nil
}

// Send godoc
func (c Mailer) Send(m *email.Email) error {
	config := c.Config.Mail
	if "" == config.Hostname {
		logger.Log("mailer.Mailer@send: invalid Hostname \"%v\"", config.Hostname)
		return nil
	}

	m.From = config.FromAddress

	addr := fmt.Sprintf("%s:%v", config.Hostname, config.Port)
	auth := smtp.PlainAuth("", config.User, config.Password, config.Hostname)

	err := m.Send(addr, auth)
	if err != nil {
		logger.Log(err)
	}

	return err
}
