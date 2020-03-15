package main

import (
	"net/http"
	"os"

	"github.com/inquizarus/godnd/magic/handlers"
	"github.com/inquizarus/godnd/magic/loaders"
	"github.com/inquizarus/gorest"
	log "github.com/sirupsen/logrus"
)

func main() {

	spellsReader, err := os.Open("data/spells.json")
	nilErrOrPanic(err)
	schoolsReader, err := os.Open("data/schools.json")
	nilErrOrPanic(err)
	spells, err := loaders.LoadSpells(spellsReader)
	nilErrOrPanic(err)
	schools, err := loaders.LoadSchools(schoolsReader)
	nilErrOrPanic(err)

	logger := log.New()
	logger.Out = os.Stdout
	logger.Formatter = &log.JSONFormatter{}
	logger.SetLevel(log.DebugLevel)
	cfg := gorest.ServeConfig{
		Port:   "8080",
		Logger: logger,
		Handlers: []gorest.Handler{
			handlers.MakeRootHandler(),
			handlers.MakeSchoolHandler(schools),
			handlers.MakeSchoolsHandler(schools),
			handlers.MakeSpellHandler(spells),
			handlers.MakeSpellsHandler(spells),
		},
		Middlewares: []gorest.Middleware{
			gorest.WithJSONContent(),
			func(f http.HandlerFunc) http.HandlerFunc {
				// Define the http.HandlerFunc
				return func(w http.ResponseWriter, r *http.Request) {
					fields := map[string]interface{}{
						"referer":         r.Referer(),
						"client-ip":       r.RemoteAddr,
						"uri":             r.URL.RequestURI(),
						"query":           r.URL.Query().Encode(),
						"x-forwarded-for": r.Header.Get("X-FORWARDED-FOR"),
					}
					logger.WithFields(fields).Info("incoming request")
					f(w, r)
				}
			},
		},
	}
	gorest.Serve(cfg)
}

func nilErrOrPanic(err error) {
	if nil != err {
		panic(err)
	}
}
