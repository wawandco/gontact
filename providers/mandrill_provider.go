package providers

import (
	"errors"
	"os"

	m "github.com/keighl/mandrill"
	"github.com/wawandco/gontact/core"
)

//SlackProvider TODO: Document me!
type MandrillProvider struct{}

//SendContact is the implementation of the SendContact function for the Mandril Provider.
func (sp MandrillProvider) SendContact(contact core.Contact) (string, error) {
	mandrillKey := os.Getenv("MANDRILL_KEY")

	if mandrillKey == "" {
		return "", errors.New("Please define your mandril key.")
	}

	client := m.ClientWithKey(mandrillKey)
	message := &m.Message{}

	message.AddRecipient(os.Getenv("MANDRILL_TO"), "", "to")
	message.FromEmail = osEnvWithDefault("MANDRILL_FROM", "gontact@wawand.co")
	message.FromName = "Gontact Mailer"
	message.Subject = osEnvWithDefault("MANDRILL_SUBJECT", "Contact")
	message.HTML = buildMessage(contact, emailTPL)

	_, err := client.MessagesSend(message)
	return "", err
}

var emailTPL = `
<p>Team,</p>
<p>
  We got a contact request from
  <em>
    {{.Name}}
  </em>
  , his/her email address is
  <em>
    {{.Email}}
  </em>
  , and he is contacting us for
  <em>
    {{.Subject}}
  </em>
</p>
<p>
  Below you can see the message he/she left in our website:
  <blockquote>
    <p>
      {{.Message}}
    </p>
  </blockquote>
</p>
<p>
  Keep in Touch!
</p>
`
