package ggb

import (
	"github.com/jonmorehouse/go-config/config"
	"log"
)

// global file queue interaction channels
var pop chan PopOperation
var push  chan PushOperation
var errorComm chan CommunicationOperation

// global variables 
var tarballCounter int64 

func ErrorHandler() {
	var operation CommunicationOperation
	// initialize our fatal errors associative arry
	fatalErrors := config.Value("FATAL_ERRORS").(map[int]bool)
	for {
		operation = <- errorComm
		// check if there is an error and if it is fatal
		if operation.err != nil && fatalErrors[operation.code] {
			log.Fatal(operation.err)
		}
	}
}

func Bootstrap() {
	// setup configuration
	envVars := []string{
		"AWS_REGION",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_ACCESS_KEY_ID",
		"TARBALL_SIZE",
		"MAX_GO_ROUTINES",
	}
	fatalErrors := map[int]bool{
		ERROR: true,
		ERROR_UPLOAD_FAILED: true,
	}
	config.New()//instantiate config package 
	err := config.Bootstrap(envVars) 
	if err != nil {
		log.Fatal(err)
	}
	err = config.Set("FATAL_ERRORS", fatalErrors)
	if err != nil {
		log.Fatal(err)
	}
	// build out global channels 
	pop = make(chan PopOperation, 1000)
	push = make(chan PushOperation, 1000)
	errorComm = make(chan CommunicationOperation, 1000)
	// start up goworker for handling errors
	go ErrorHandler()
}



