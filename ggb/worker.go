package ggb

import (
	"sync"
	"sync/atomic"
)

func Worker(waitGroup * sync.WaitGroup, commChannel chan CommunicationOperation) {

	defer waitGroup.Done()
	var tarball * Tarball
	finished := false 
	popResponseChannel := make(chan PopResponseOperation)
	// add file to the current tarball. If its big enough, upload it and reset pointer
	addFile := func(file * File) {
		if tarball == nil {
			// increase tarball counter by 1
			id := atomic.AddInt64(&tarballCounter, 1)	
			// create a new tarball
			tarball = NewTarball(id)
		}
		// add a file to the tarball now that we know it exists successfully
		err := tarball.AddFile(file)
		if err != nil {
			commChannel <- CommunicationOperation{code: ERROR_TARBALL_FILE, err: err}
		}
		if tarball.Full {
			tarball.Upload()
			// reset the tarball as we need to generate a new one 
			tarball = nil
		}
	}
	// need a channel for grabbing file queue elements as needed
	for {
		select {
		case _ = <- commChannel:
			// in case somethign pushes into our communication channel to force a stop -- this can be removed later?
			tarball.Upload()	
			finished = true
		default: 
			// try to make a pop from teh fileQueue as needed
			pop <- PopOperation{channel: popResponseChannel}
			// now wait for a response on this element
			response := <- popResponseChannel
			// no files left in queue
			if response.err != nil {
				tarball.Upload()	
				finished = true
			} else { // we got a file successfully - add it and process it accordingly
				addFile(response.file)// add the response and process the tarball accordingly
			}
		}
		if finished {
			break
		}
	}
}


