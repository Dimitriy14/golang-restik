package main

import (
	"flag"
	"log"
	"net/http"

	apploader "github.com/Dimitriy14/golang-restik/src/app-loader"
	"github.com/Dimitriy14/golang-restik/src/config"
	"github.com/Dimitriy14/golang-restik/src/logger"
)

func main() {
	configPath := flag.String("-conf", "../config.json", "-config ")

	if *configPath != "" {
		config.FilePath = *configPath
	}

	err := apploader.LoadApplicationServices()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr: config.Conf.ListenURL,
		//Handler: there should be router,
	}

	logger.Log.Infof("tx", "Started serving at: %s", config.Conf.ListenURL)
	if err := server.ListenAndServe(); err != nil {
		logger.Log.Error("", "", "==== Restik stopped due to error: %v", err)
	}
}
