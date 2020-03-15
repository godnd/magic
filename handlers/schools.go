package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/inquizarus/godnd/magic/converters"
	"github.com/inquizarus/godnd/magic/models"
	"github.com/inquizarus/gorest"
)

// MakeSchoolsHandler creates handler for the "/schools" path
func MakeSchoolsHandler(schools []models.School) *gorest.BaseHandler {
	return &gorest.BaseHandler{
		Name: "SchoolsHandler",
		Path: "/schools",
		Get: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			defer r.Body.Close()
			bbytes, _ := json.Marshal(schools)
			w.Write(bbytes)
		},
	}
}

// MakeSchoolHandler creates handler for the "/schools/{name}" path
func MakeSchoolHandler(schools []models.School) *gorest.BaseHandler {
	return &gorest.BaseHandler{
		Name: "SchoolHandler",
		Path: "/schools/{name}",
		Get: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			name := converters.KebabCaseStringToRegular(p["name"])
			defer r.Body.Close()
			for _, school := range schools {
				if school.Name == name {
					bbytes, _ := json.Marshal(school)
					w.Write(bbytes)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		},
	}
}
