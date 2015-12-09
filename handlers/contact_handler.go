package handlers

import "net/http"

//ContactHandler handles requests done to our /contact endpoint and pass our authentication method.
func ContactHandler(http.ResponseWriter, *http.Request) {
	//1. Validate the request
	//2. Respond 422 if not complete
	//3. Process if complete
	//4. Respond with 201 if provider responds OK
	//5. Respond with 500 is profider responds with error
}
