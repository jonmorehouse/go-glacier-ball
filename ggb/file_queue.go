package ggb

import (

	"container/heap"
	"fmt"
	"sync"
)

type FileQueue []*File

func NewFileQueue() * FileQueue {

	fq := &FileQueue{}

	heap.Init(fq)

	return fq
}

func (fq FileQueue) Len() int {

	return len(fq)

}

/*
func (fq FileQueue) Push(file * File) {

	heap.Push(fq, file)

}

func (fq FileQueue) Pop() * File {

	if !fq.Empty() {

		return heap.Pop(fq)		

	} else {

		return nil
	}

}

func (fq FileQueue) Empty() bool { 

	if len(fq) == 0 {

		return true

	} else {

		return false

	}
}
*/

/*
	assumes that all file structs are for legitimate files 
	assumes it will be run as a go routine and will be told to stop by the parent 
	communication channel = read/write
	push channel = read 
	pop channel = read -- this is really a request pop

*/
func FileQueueManager(waitGroup * sync.WaitGroup, pushChannel chan PushOperation, popChannel chan PopOperation, communicationChannel chan CommunicationOperation) {

	// initialize queue - this should be a pointer
	queue := NewFileQueue()

	// initialize queue worker 
	waitGroup.Add(1)
	finished := false

	// now lets go ahead 
	for {

		select {

		// step 1 - see if we have anything for communication
		// this can be a read or write operation
		case push := <- pushChannel:

			// push into the channel  	
			fmt.Println(push.message)

		// step 2 - queue up any files that need to be queued
		case pop := <- popChannel:

			fmt.Println(pop.message)

		// step 3 - respond to any pop requests as needed
		case comm := <- communicationChannel:

			if comm.code == ALL_JOBS_SUBMITTED {

				finished = true
			}
		}

		// this worker is finished
		if finished /*&& list.Empty()*/ {

			waitGroup.Done()
			break

		// end in the case of an error that was communicated from an outside process
		} else if finished /*&& systemError */{

			waitGroup.Done()
			break
		}
	}


}





