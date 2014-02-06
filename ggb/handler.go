package ggb

import (

)

// declare constants as needed
const GB int64 = 1024 * 1024 * 1024 // 1 GB

// now  create queue sizes as needed
const (
	
	SMALL_QUEUE = 1 * GB // maximum size for small queue
	LARGE_QUEUE = 2 * GB // maximum size for large queue
)

const (

	WORKER_BUSY = 0 // ie its uploading
	WORKER_AVAILABLE = 1 // upload finished
	WORKER_REJECT_PATH = 2 // expect byte string with the path that was rejected - this file needs to be reprocessed

	PROCESSOR_SUBMIT_PATH = 3 //
	PROCESSOR_FINISHED = 4 // all pathes have been processed
)

// handler is the struct that is responsible for dealing with input queue etc
type Handler struct {

	// workers -- go-routines for workers
	// queues -- large / small queues
}

// go routine that is responsible for managing 
func HandlerWorker() {

	// manages the busy / available workers 
	// manages a spit back file -- puts this back in the normal queue
	// this shares the HAndler struct between the processor and itself
}

func AddFile (path string) {



}

func AddFiles (paths[]string) {




}


