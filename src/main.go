package main

import (
	"fmt"

	"net/http"
	"os"

	//"docs/docs.go"
	//"github.com/swaggo/http-swagger" // http-swagger middleware

	// http-swagger middleware
	src "github.com/dhf0820/caDbConnector/src"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	//m "pkg/model"
)

// http-swagger middleware

// @title FHIR API for ChartArchive
// @version .2
// @description fhirgo creates an interface to various FHIR systems. initially it is Cerner.
// There is no security at this time so only fake PHI is being used.
// Supports patient/encounters/documents(Diagnostic Reports)

//var (
// 	config  *Config
// 	name    string
// 	mode    string
// 	port    string
// 	mainErr error
// )

func main() {
	fmt.Println("Start Connector")
	godotenv.Load()

	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true

	log.SetFormatter(Formatter)
	log.SetLevel(log.DebugLevel)
	log.Infof("Getting MONGODB url")
	mongoURL, err := os.LookupEnv("MONGODB")
	if err != true {
		log.Errorf("Error in MONGODB lookup: %v", err)
		log.Fatal(err)
	}

	log.Infof("Calling Initialize all")
	config := src.InitializeAll(mongoURL, "")

	port := config.Port()
	//mode := config.Mode()
	router := src.NewRouter
	config.SetRouter(router)
	// log.Infof("Serving %s FHIR interface VERSION %s in %s mode on port: %s", config.Source(), config.ServerVersion(), mode, port)
	// log.Infof("URL to use: %s", config.BaseUrl())
	// log.Infof("Image URL: %s", config.ImageURL())
	// log.Infof("Ca Image URL: %s", config.Env("caImageURL"))
	// log.Infof("Ca Server url: %s", config.Env("caServerURL"))
	mainErr := http.ListenAndServe(port, router)
	if mainErr != nil {
		log.Errorf("Main error: %v", mainErr)
	}

}
