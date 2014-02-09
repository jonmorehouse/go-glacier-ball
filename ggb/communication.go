package ggb

// 
const statusCodes (

	// establish various status code elements as needed
	REQUEST_QUEUE_STATUS = 00
	QUEUE_EMPTY = 01
	ALL_JOBS_SUBMITTED = 02

	// error codes for various elements as needed
	ERROR_FILE_NOT_FOUND = 10
	ERROR_UPLOAD_FAILED = 11

	// 
)

struct ErrorOperation {

	// error code


}

struct StatusOperation {

	// transferring of statuses for input/output from workers/queues etc

}

struct RequestOperation {

	// request a file. This should respond with a status code

}

struct QueueOperation {

	// 

}



