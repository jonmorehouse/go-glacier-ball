package ggb

import (
	"github.com/jonmorehouse/go-config/config"
	"log"
	"fmt"
)

// global file queue interaction channels
var pop chan PopOperation
var push  chan PushOperation
var errorComm chan CommunicationOperation

func ErrorHandler() {

	var operation CommunicationOperation

	for {
		operation = <- errorComm
		
		fmt.Println(operation)
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
	config.New()//instantiate config package 
	err := config.Bootstrap(envVars) 
	if err != nil {
		log.Fatal(err)
	}
	// build out global channels 
	pop = make(chan PopOperation, 1000)
	push = make(chan PushOperation, 1000)
	errorComm = make(chan CommunicationOperation, 1000)
}
