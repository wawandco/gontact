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
}

//ContactHandler handles requests done to our /contact endpoint and pass our authentication method.
func ContactHandler(rw http.ResponseWriter, req *http.Request) {

	contact := core.Contact{}
	decoder := schema.NewDecoder()
	decoder.Decode(&contact, req.Form)

	_, err := validator.ValidateStruct(contact)
	if err != nil {
		rw.WriteHeader(422)
		rw.Write([]byte(err.Error()))
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
