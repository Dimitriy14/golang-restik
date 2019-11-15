package main

import (
	"flag"
	"log"
	"net/http"

	apploader "github.com/Dimitriy14/golang-restik/app-loader"
	"github.com/Dimitriy14/golang-restik/config"
	"github.com/Dimitriy14/golang-restik/logger"
	"github.com/Dimitriy14/golang-restik/services"
	"github.com/urfave/negroni"
)

func main() {
	configPath := flag.String("-conf", "config.json", "-config ")

	if *configPath != "" {
		config.FilePath = *configPath
	}

	err := apploader.LoadApplicationServices()
	if err != nil {
		log.Fatal(err)
	}

	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	negroniLogger := negroni.NewLogger()
	negroniLogger.ALogger = logger.NewNegroniLogger(logger.Log)

	middlewareManager.Use(negroniLogger)
	middlewareManager.UseHandler(services.NewRouter())

	server := &http.Server{
		Addr:    config.Conf.ListenURL,
		Handler: middlewareManager,
	}

	logger.Log.Infof("", "Started serving at: %s", config.Conf.ListenURL)
	if err := server.ListenAndServe(); err != nil {
		logger.Log.Error("", "", "==== Restik stopped due to error: %v", err)
	}
}
