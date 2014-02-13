package ggb

import (
	"sync"
	"fmt"
	"github.com/jonmorehouse/go-config/config"
	"math"
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

func ProcessorManager(waitGroup * sync.WaitGroup, filePaths * []string) {

	// localWaitGroup := sync.Waitgroup
	numberWorkers := config.Value("MAX_GO_ROUTINES").(int) / 2
	filesPerWorker := int(math.Ceil(float64(float64((len(*filePaths) + 3))/float64(numberWorkers))))
	output := []string{}
	for i := 0; i < numberWorkers; i++ {
		leftIndex := i * filesPerWorker
		rightIndex := ((i+1)*filesPerWorker)
		// make sure we don't overshoot the right side of the array and add in empty values
		if rightIndex > len(*filePaths) {
			rightIndex = len(*filePaths)
		}
		workerFilePaths := copy(*filePaths)[leftIndex:rightIndex]
	}

	fmt.Println(len(output))

}

