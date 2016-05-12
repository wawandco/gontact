package providers

import (
	"errors"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/wawandco/gontact/core"
)

//SendgridProvider
type SendgridProvider struct{}

//SendContact is the implementation of the SendContact function for the Mandril Provider.
func (sp SendgridProvider) SendContact(contact core.Contact) (string, error) {
	sendgridKey := os.Getenv("SENDGRID_KEY")

	if sendgridKey == "" {
		return "", errors.New("Please define your sendgrid key.")
	}

	sg := sendgrid.NewSendGridClientWithApiKey(sendgridKey)

	message := sendgrid.NewMail()
	message.AddTo(os.Getenv("MAIL_TO"))
	message.SetFrom(osEnvWithDefault("MAIL_FROM", "gontact@wawand.co"))
	message.SetSubject(osEnvWithDefault("MAIL_SUBJECT", "Contact"))
	message.SetHTML(buildMessage(contact, emailTPL))
	err := sg.Send(message)

	return "", err
}
