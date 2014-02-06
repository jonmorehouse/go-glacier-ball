//Handler is responsible for creating / managing go routines that are working on the current queue
// this should be its own go routine - so we can spin up 1 / 2 of these
package ggb

import (

	//"file"
)

const GB int64 = 1024 * 1024 * 1024 // 1 GB

// now  create queue sizes as needed
const (
	
	SMALL_QUEUE = 1 * GB // maximum size for small queue
	LARGE_QUEUE = 2 * GB // maximum size for large queue
)

// handler is the struct that is responsible for dealing with input queue etc
type Handler struct {

	// workers -- go-routines for workers
	// queues -- large / small queues

}







