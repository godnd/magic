package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/inquizarus/godnd/magic/converters"
	"github.com/inquizarus/godnd/magic/models"
	"github.com/inquizarus/gorest"
)

// MakeSpellsHandler creates handler for the "/spells" path
func MakeSpellsHandler(spells []models.Spell) *gorest.BaseHandler {
	return &gorest.BaseHandler{
		Name: "spells",
		Path: "/spells",
		Get: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			defer r.Body.Close()
			bbytes, _ := json.Marshal(spells)
			w.Write(bbytes)
		},
	}
}

// MakeSpellHandler creates handler for the "/spells/{name}" path
func MakeSpellHandler(spells []models.Spell) *gorest.BaseHandler {
	return &gorest.BaseHandler{
		Name: "spell",
		Path: "/spells/{name}",
		Get: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			name := converters.KebabCaseStringToRegular(p["name"])
			defer r.Body.Close()
			for _, spell := range spells {
				if spell.Name == name {
					bbytes, _ := json.Marshal(spell)
					w.Write(bbytes)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		},
	}
}
