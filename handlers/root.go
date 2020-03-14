package handlers

import (
	"net/http"

	"github.com/inquizarus/gorest"
)

// MakeRootHandler creates handler for the "/" path
func MakeRootHandler() *gorest.BaseHandler {
	return &gorest.BaseHandler{
		Name: "RootHandler",
		Path: "/",
		Get: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			defer r.Body.Close()
			w.Write([]byte(`{"message": "spelldata?"}`))
		},
	}
}
