package ggb

import (
	"sync"
	"github.com/jonmorehouse/go-config/config"
)

func Worker(waitGroup * sync.WaitGroup, commChannel chan CommunicationOperation) {

	defer waitGroup.Done()
	tarball, err := NewTarball(config.Value("TARBALL_PREFIX").(string))
	if err != nil {
		commChannel <- CommunicationOperation{err: err}
	}
	finished := false
	popReciever := make(chan PopResponseOperation, 5)
	// finalize tarball as needed etc
	closeTarball := func() {
		finished = true
		// try to upload the tarball 
		if err := tarball.Upload(); err != nil {
			commChannel <- CommunicationOperation{err: err}
		}
	}
	for {
		select {
		case _ = <- commChannel:
			closeTarball()
		default://pop a new file
			// submit our reciever channel to the pop as needed
			pop <- PopOperation{channel: popReciever}
			response := <- popReciever
			if response.err != nil {
				closeTarball()
			} else {
				newTarball, err := tarball.AddFile(response.file)
				if err != nil || newTarball == nil {
					finished = true
					commChannel <- CommunicationOperation{err: err}
				}
				tarball = newTarball
			}
		}
		if finished {
			break
		}
	}
}

func WorkerManager(wg * sync.WaitGroup) {

	defer wg.Done()

	//var lwg sync.WaitGroup
	// check the number / size of file and create accordingly



}

