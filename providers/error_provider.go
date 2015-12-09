package providers

import (
	"errors"

	"github.com/wawandco/gontact/core"
)

//InternalProvider is a provider written to be used on the testing.
type ErrorProvider struct{}

//SendContact Implements the Provider interface, this only returns an error! :D.
func (ip ErrorProvider) SendContact(con core.Contact) (string, error) {
	return "This is an error provider response.", errors.New("Simple Error Response!.")
}
