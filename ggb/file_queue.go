package ggb

import (
	"container/list"
	"sync"
	"errors"
)

/*
	assumes that all file structs are for legitimate files 
	assumes it will be run as a go routine and will be told to stop by the parent 
	communication channel = read/write
	push channel = read 
	pop channel = read -- this is really a request pop
*/
func FileQueueManager(waitGroup * sync.WaitGroup, communicationChannel chan CommunicationOperation) {

	// initialize queue - this should be a pointer
	queue := list.New()
	finished := false
	errorReported := false
	waitGroup.Add(1)

	for {
		select {
		// step 1 - see if we have anything for communication
		case pushOperation := <- push:
			// push into the channel  	
			queue.PushBack(pushOperation.file)
		// step 2 - queue up any files that need to be queued
		case popOperation := <- pop:
			if queue.Len() == 0 {//no elements to pass back -- pas an error
				popOperation.channel <- PopResponseOperation{err: errors.New("Queue empty")}
			} else {
				// now lets grab the last element in the list
				element := queue.Back()
				// grab the actual file 
				file := element.Value.(*File)
				// now remove the file from the list  
				queue.Remove(element)
				// now lets pipe the file pointer to the file as needed
				popOperation.channel <- PopResponseOperation{file: file}
			}
		// step 3 - respond to any pop requests as needed
		case comm := <- communicationChannel:
			if comm.code == ALL_JOBS_SUBMITTED {
				finished = true
			} else if comm.code == ERROR {
				errorReported = true
			} else if comm.code == QUEUE_STATUS {
				// create a response
				response := CommunicationOperation{code: QUEUE_STATUS_RESPONSE, message: queue.Len()}
				// push the response structure to the channel that is expecting it
				comm.channel <- response
			}
		}// end of select statement 

		// this worker is finished
		if finished && queue.Len() == 0 {

			// pass a message to the communication manager 
			break

		// end in the case of an error that was communicated from an outside process
		} else if errorReported {

			break
		}
	}
	waitGroup.Done()
}

