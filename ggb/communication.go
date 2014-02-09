package ggb

const (

	// establish various status code elements as needed
	REQUEST_QUEUE_STATUS = 00
	QUEUE_EMPTY = 01
	ALL_JOBS_SUBMITTED = 02

	// error codes for various elements as needed
	ERROR_FILE_NOT_FOUND = 10
	ERROR_UPLOAD_FAILED = 11

	// large / small file sizes
	LARGE_FILE = 20
	SMALL_FILE = 21

)

type CommunicationOperation struct {

	// transferring of statuses for input/output from workers/queues etc
	err error
	code int 

	message string
}

type PopOperation struct {

	// request a file. This should respond with a status code
	// should pass a pointer to a file channel
	// this should be writable on the the worker's end
	file chan * File // this is where we will spit the file back
	fileType int 

	message string
}

type PushOperation struct {

	// assumes a legitimate file being passed in as needed
	// file pointer that needs to be passed in 	
	file * File

	message string
	
	// status type
	fileType int

}


