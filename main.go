package main

import (
	"os"

	"github.com/inquizarus/godnd/spells/handlers"
	"github.com/inquizarus/gorest"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.Out = os.Stdout
	logger.Formatter = &log.JSONFormatter{}
	cfg := gorest.ServeConfig{
		Port:   "8080",
		Logger: logger,
		Handlers: []gorest.Handler{
			handlers.MakeRootHandler(),
		},
		Middlewares: []gorest.Middleware{
			gorest.WithJSONContent(),
		},
	}
	gorest.Serve(cfg)
}
