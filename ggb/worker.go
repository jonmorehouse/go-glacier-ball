package ggb

import (
	"sync"
	//"sync/atomic"
	//"github.com/jonmorehouse/go-config/config"
	"fmt"
)


func Worker(waitGroup * sync.WaitGroup, commChannel chan CommunicationOperation) {

	defer waitGroup.Done()
	//tarball := NewTarball()
	finished := false
	popReciever := make(chan PopResponseOperation, 5)
	for {

		select {
		case _ = <- commChannel:
			// handle communication here

		default://pop a new file
			// submit our reciever channel to the pop as needed
			pop <- PopOperation{channel: popReciever}
			response := <- popReciever

			if response.err != nil {
				finished = true
			} else {

				fmt.Println(response.file)
			}
		}

		if finished {
			break
		}
	}
}


