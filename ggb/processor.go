package ggb

import (
	"sync"
	"github.com/jonmorehouse/go-config/config"
	"math"
)

/*
	Iterate through a list of files

	create file structs for each and submit them to our file manager
	errors are reported in the communication channel
*/
func Processor(waitGroup * sync.WaitGroup, comm chan CommunicationOperation, files []string) {
	waitGroup.Add(1)
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

func ProcessorManager(waitGroup * sync.WaitGroup, filePaths * []string) {
	var localWaitGroup sync.WaitGroup
	waitGroup.Add(1)
	numberWorkers := config.Value("MAX_GO_ROUTINES").(int) / 2
	filesPerWorker := int(math.Ceil(float64(float64((len(*filePaths) + 3))/float64(numberWorkers))))
	for i := 0; i < numberWorkers; i++ {
		// generate how long the individual worker's path array is
		index := filesPerWorker
		if index > len(*filePaths) {
			index = len(*filePaths)
		}
		// generate the filePaths slice and then pass it to the go routine that will be processing this
		workerFilePaths := make([]string, index)
		copy(workerFilePaths, (*filePaths)[0:index])
		go Processor(&localWaitGroup, errorComm, workerFilePaths)
		// remove this piece of the slice from the original slice 
		(*filePaths) = (*filePaths)[index:]
	}
	localWaitGroup.Wait()
	waitGroup.Done()
}


