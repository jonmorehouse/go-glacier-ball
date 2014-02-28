package ggb

import (
	"strings"
	"bufio"
	"io"
	"sync"
)

func FileUploader(wg * sync.WaitGroup, input io.Reader) {
		
	defer wg.Done()
	reader := bufio.NewReader(input)
	filePaths := []string{}
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil {
			break
		} else {
			// now we want to pipe this to the processor worker as needed
			// this functionality can come later
			filePaths = append(filePaths, line)
		}
	}
	//fileCount := len(filePaths)
	// we now have all filePaths as strings - begin processing each file
	var fwg sync.WaitGroup
	fwg.Add(1)
	fcomm := make(chan CommunicationOperation)
	go FileQueueManager(&fwg, fcomm)
	// now load up and process all the files as needed
	var pwg sync.WaitGroup
	pwg.Add(1)
	go ProcessorManager(&pwg, &filePaths)
	pwg.Wait()
	// now go ahead and start up workers as needed
	var wwg sync.WaitGroup	
	wcomm := make(chan CommunicationOperation)
	wwg.Add(1)
	go Worker(&wwg, wcomm)
}

func TextUploader(wg sync.WaitGroup, input io.Reader) {


}


