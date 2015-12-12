package providers

import "github.com/wawandco/Gontact/Godeps/_workspace/src/github.com/wawandco/gontact/core"

//InternalProvider is a provider written to be used on the testing.
type InternalProvider struct{}

var InternalProviderCount = 0

//SendContact Implements the Provider interface.
func (ip InternalProvider) SendContact(con core.Contact) (string, error) {
	InternalProviderCount = InternalProviderCount + 1
	return "Ok", nil
}
