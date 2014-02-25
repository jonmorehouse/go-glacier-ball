package ggb

import (
	"sync"
	"os"
	"bufio"
	"fmt"
	"github.com/jonmorehouse/go-config/config"
)

func GGB() {
	Bootstrap()
	// initialize the filepaths that need to be processed
	reader := bufio.NewReader(os.Stdin)
	filePaths := []string{}
	for {
		path, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		filePaths = append(filePaths, path)
	}

	if len(filePaths) == 0 {
		return
	}

	files := len(filePaths)

	// set up the workers that will process all of the uploads
	comm := make(chan CommunicationOperation)
	var wg sync.WaitGroup
	var pwg sync.WaitGroup

	// start file queue manager
	go FileQueueManager(&wg, comm)
	// process all files as needed
	pwg.Add(1)
	ProcessorManager(&pwg, &filePaths)
	pwg.Wait()
	// once the processor manager is finished, all files have been submitted
	// let the processor know
	comm <- CommunicationOperation{code: ALL_JOBS_SUBMITTED}
	
	for i := 0; i < files; i++ {

				
	}


	return
	// now lets start up 1 worker (for now)
	wg.Add(1)
	go Worker(&wg, comm)
	wg.Wait()	

	// now print out the prefix for these files
	fmt.Println(config.Value("TARBALL_PREFIX").(string))
}

