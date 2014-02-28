package ggb

import (
	"io"
	"sync"
)

func FileUploader(wg * sync.WaitGroup, input io.Reader) {

	defer wg.Done()
	fileComm := make(chan CommunicationOperation)
	// start up the filequeue
	go FileQueueManager(&sync.WaitGroup{}, fileComm)
	// now we need to process all the files as needed
	var lwg sync.WaitGroup	
	lwg.Add(1)
	go ProcessorManager(&lwg, input)
	lwg.Wait()

}

func TextUploader(wg sync.WaitGroup, input io.Reader) {


}


