package providers

import (
	"errors"
	"os"

	m "github.com/keighl/mandrill"
	"github.com/wawandco/gontact/core"
)

//SendgridProvider
type SendgridProvider struct{}

//SendContact is the implementation of the SendContact function for the Mandril Provider.
func (sp SendgridProvider) SendContact(contact core.Contact) (string, error) {
	mandrillKey := os.Getenv("SENDGRID_KEY")

	if mandrillKey == "" {
		return "", errors.New("Please define your mandril key.")
	}

	client := m.ClientWithKey(mandrillKey)
	message := &m.Message{}

	message.AddRecipient(os.Getenv("MAIL_TO"), "", "to")
	message.FromEmail = osEnvWithDefault("MAIL_FROM", "gontact@wawand.co")
	message.FromName = "Gontact Mailer"
	message.Subject = osEnvWithDefault("MAIL_SUBJECT", "Contact")
	message.HTML = buildMessage(contact, emailTPL)

	_, err := client.MessagesSend(message)
	return "", err
}
