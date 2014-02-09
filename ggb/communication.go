package ggb

const (

	// establish various status code elements as needed
	QUEUE_STATUS = 00
	QUEUE_EMPTY = 01
	ALL_JOBS_SUBMITTED = 02

	// error codes for various elements as needed
	ERROR = 10 
	ERROR_FILE = 11
	ERROR_UPLOAD_FAILED = 12

	// large / small file sizes
	LARGE_FILE = 20
	SMALL_FILE = 21
)

type CommunicationOperation struct {

	// transferring of statuses for input/output from workers/queues etc
	err error
	code int 
	message interface{}

	// incase we want to pass a message back 
	channel chan CommunicationOperation 
}

type PopOperation struct {

	// request a file. This should respond with a status code
	// should pass a pointer to a file channel
	// this should be writable on the the worker's end
	channel chan * File // this is where we will spit the file back
}

type PushOperation struct {

	// assumes a legitimate file being passed in as needed
	// file pointer that needs to be passed in 	
	file * File

}


