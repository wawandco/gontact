package providers

import "github.com/wawandco/gontact/core"

//Provider is a interface build to set a baseline for implementations of different
//transactional email providers.
type Provider interface {
	SendContact(contact core.Contact) (string, error)
}
