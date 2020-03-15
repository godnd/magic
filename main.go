package main

import (
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
		},
	}
	gorest.Serve(cfg)
}

func nilErrOrPanic(err error) {
	if nil != err {
		panic(err)
	}
}
