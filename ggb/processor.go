package ggb

import (

	"sync"
)

/*
	Iterate through a list of files

	create file structs for each and submit them to our file manager
	errors are reported in the communication channel
*/
func Processor(waitGroup * sync.WaitGroup, push chan PushOperation, comm chan CommunicationOperation, files []string) {

	waitGroup.Add(1)

	// for each file
	// create a new file struct and initialize the elements
	// pass the pointer to the workerQueue as needed
	for i := range files {

		// create the file as needed
		file, err := NewFile(files[i])

		if err != nil {

			// pass the error to our error handler
			comm <- CommunicationOperation{code: ERROR_FILE, message: &err}

		} else {

			// push this file to the push channel as needed
			push <- PushOperation{file: &file}
		}
	}

	waitGroup.Done()
}

