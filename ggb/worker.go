package ggb

import (

	"sync"
)


/*
	args:
		1.) wait group
		2.) communication channel (for error communication)
		3.) push channel - for rejecting files that aren't necessary
		4.) pop channel - for requesting files too add 


*/
func Worker(waitGroup * sync.WaitGroup, communication chan CommunicationOperation, pop chan PopOperation, push chan PushOperation) {

	//waitGroup.Add(1)

	// should recieve a notice when this is the last run
	// this should be when the queue is empty
	

	//waitGroup.Done(1)

}
