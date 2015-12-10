package providers

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"text/template"

	"github.com/wawandco/gontact/core"
)

//SlackProvider TODO: Document me!
type SlackProvider struct{}

//SendContact is the implementation of the SendContact function for the Slack Provider.
func (sp SlackProvider) SendContact(contact core.Contact) (string, error) {

	message := buildSlackMessage(contact)
	hookURL := osEnvWithDefault("SLACK_WEBHOOK_URL", "")

	if hookURL == "" {
		return "", errors.New("You need to define SLACK_WEBHOOK_URL on your Gontact server")
	}

	channel := osEnvWithDefault("SLACK_CHANNEL", "notifications")
	username := osEnvWithDefault("SLACK_USERNAME", "Gontact")
	emoji := osEnvWithDefault("SLACK_EMOJI", ":mailbox:")

	payload := "{\"channel\": \"" + channel + "\", \"username\": \"" + username + "\", \"text\": \"" + message + "\", \"icon_emoji\": \":" + emoji + ":\"}"

	body := bytes.NewBufferString(payload)
	response, err := http.Post(hookURL, "text/plain", body)

	if err != nil {
		return "", err
	} else if response.StatusCode == 500 {
		return "Error", errors.New("Server error")
	}

	return "OK", nil
}

var slackMessageString = `Received a Website Contact \n *Name*: {{.Name}} \n *Email*: {{.Email}} \n *Message*: {{.Message}}`

func buildSlackMessage(contact core.Contact) string {
	messageTPL, _ := template.New("slack.message").Parse(slackMessageString)
	var doc bytes.Buffer
	messageTPL.Execute(&doc, contact)
	message := doc.String()
	return message
}

func osEnvWithDefault(name string, def string) string {
	value := os.Getenv(name)
	if value == "" {
		value = def
	}

	return value
}
