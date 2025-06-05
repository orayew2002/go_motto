package main

import (
	"github.com/orayew2002/go_motto/internal/app"
	"github.com/orayew2002/go_motto/internal/domains"
	"github.com/orayew2002/go_motto/pkg/log"
)

func main() {
	logger := log.InitSlog(true, "logs/app.log")

	app.Run(domains.AppDependencies{
		Logger: logger,
	})
}
