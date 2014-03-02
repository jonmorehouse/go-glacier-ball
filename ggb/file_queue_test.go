package ggb

import (
	. "launchpad.net/gocheck"
	//"fmt"
)

type FileQueueSuite struct {
	GGBSuite
}

// bootstrap the suite into the current environment
var _ = Suite(&FileQueueSuite{})

// now lets test the basic functionality of the different elements as needed
func (s * FileQueueSuite) SetUpTest(c *C) {
	
	s.Bootstrap()
	s.comm = make(chan CommunicationOperation, 2)
	// now we can start the go routine as needed
	go FileQueueManager(&s.wg, s.comm)
	// push all the files into the shared queue
	for i := range s.files {
		push <- PushOperation{file: s.files[i]}
	}
}

func (s *FileQueueSuite) TearDownTest(c *C) {
	s.Breakdown()
}

/*
func (s *FileQueueSuite) TestFilePush(c *C) {
	// push a single file into the queue
	push <- PushOperation{file: s.files[0]}
	// now we need to test the length of the element as needed
	request := CommunicationOperation{code: QUEUE_TOTAL_FILES, channel: make(chan CommunicationOperation, 2)}
	s.comm <- request
	response := <- request.channel
}

func (s *FileQueueSuite) TestFilePopping(c *C) {
	reciever := make(chan PopResponseOperation, 10)
	popOperation := PopOperation{channel: reciever}
	// now lets make a few pop requests
	for _ = range s.files {
		pop <- popOperation
		popResponse := <- reciever
		c.Assert(popResponse.err, IsNil)
		c.Assert(popResponse.file, NotNil)
	}
}
*/

func (s *FileQueueSuite) TestFileMessaging(c *C) {
	
	testOperation := func(code int) CommunicationOperation {
		request := CommunicationOperation{channel: make(chan CommunicationOperation), code: code}
		s.comm <- request
		response := <- request.channel
		return response
	}

	push <- PushOperation{file: s.files[0], channel: make(chan CommunicationOperation, 3)}
	push <- PushOperation{file: s.files[0], channel: make(chan CommunicationOperation, 3)}
	push <- PushOperation{file: s.files[0], channel: make(chan CommunicationOperation, 3)}
	// there is no guarantee of which channels will take priority
	// ie: push could be processed before comm or vice versa 
	response := testOperation(QUEUE_TOTAL_FILES)
	//c.Assert(response, NotNil)
	//fmt.Println(response.message)
	//c.Assert(response.message, Equals, 6)
}





