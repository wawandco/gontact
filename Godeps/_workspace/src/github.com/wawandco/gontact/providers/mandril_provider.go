package providers

import (
	"errors"
	"os"

	m "github.com/wawandco/Gontact/Godeps/_workspace/src/github.com/keighl/mandrill"
	"github.com/wawandco/Gontact/Godeps/_workspace/src/github.com/wawandco/gontact/core"
)

//SlackProvider TODO: Document me!
type MandrilProvider struct{}

//SendContact is the implementation of the SendContact function for the Mandril Provider.
func (sp MandrilProvider) SendContact(contact core.Contact) (string, error) {
	mandrilKey := os.Getenv("GONTACT_MANDRIL_KEY")

	if mandrilKey == "" {
		return "", errors.New("Please define your mandril key.")
	}

	client := m.ClientWithKey(mandrilKey)
	message := &m.Message{}

	message.AddRecipient("bob@example.com", "Bob Johnson", "to")
	message.FromEmail = osEnvWithDefault("MANDRIL_FROM", "gontact@wawand.co")
	message.FromName = "Gontact Mailer"
	message.Subject = osEnvWithDefault("MANDRIL_SUBJECT", "Contact")
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
