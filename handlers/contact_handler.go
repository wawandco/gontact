package handlers

import (
	"net/http"
	"os"

	validator "github.com/wawandco/gontact/Godeps/_workspace/src/github.com/asaskevich/govalidator"
	"github.com/wawandco/gontact/Godeps/_workspace/src/github.com/gorilla/schema"
	"github.com/wawandco/gontact/core"
	"github.com/wawandco/gontact/providers"
)

var registeredProviders = map[string]providers.Provider{
	"INTERNAL": providers.InternalProvider{},
	"ERROR":    providers.ErrorProvider{},
	"SLACK":    providers.SlackProvider{},
}

//ContactHandler handles requests done to our /contact endpoint and pass our authentication method.
func ContactHandler(rw http.ResponseWriter, req *http.Request) {
	contact, _ := parseContact(req)

	_, err := validator.ValidateStruct(contact)
	if err != nil {
		rw.WriteHeader(422)
		rw.Write([]byte(err.Error()))
		return
	}

	providerName := os.Getenv("GONTACT_PROVIDER")
	provider := registeredProviders[providerName]

	if provider != nil {
		_, err = provider.SendContact(contact)

		if err != nil {
			rw.WriteHeader(500)
			return
		}
	}

	rw.WriteHeader(201)
}

func parseContact(req *http.Request) (core.Contact, error) {
	err := req.ParseForm()

	if err != nil {
		return core.Contact{}, err
	}

	contact := core.Contact{}
	decoder := schema.NewDecoder()
	decoder.Decode(&contact, req.PostForm)

	return contact, nil
}
