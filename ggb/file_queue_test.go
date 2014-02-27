package ggb

import (
	"sync"
	. "launchpad.net/gocheck"
)

type FileQueueSuite struct {
	communication chan CommunicationOperation
	waitGroup sync.WaitGroup
	files []*File
	filePaths []string
}

// bootstrap the suite into the current environment
var _ = Suite(&FileQueueSuite{})

// now lets test the basic functionality of the different elements as needed
func (s * FileQueueSuite) SetUpTest(c *C) {
	s.files = CreateFileList(5)
	for i := range s.files {
		s.filePaths = append(s.filePaths, s.files[i].path)
	}
	s.communication = make(chan CommunicationOperation)
	// now we can start the go routine as needed
	go FileQueueManager(&s.waitGroup, s.communication)
	// push all the files into the shared queue
	for i := range s.files {
		push <- PushOperation{file: s.files[i]}
	}
}

func (s *FileQueueSuite) TearDownTest(c *C) {
	RemoveFileList(&s.files)
	s.filePaths = []string{}
}

func (s *FileQueueSuite) TestFilePush(c *C) {
	// push a single file into the queue
	push <- PushOperation{file: s.files[0]}
	// now we need to test the length of the element as needed
	request := CommunicationOperation{code: QUEUE_STATUS, channel: make(chan CommunicationOperation, 2)}
	s.communication <- request
	response := <- request.channel
	c.Assert(response.message.(int), Equals, 6)
}

func (s *FileQueueSuite) TestQueueCommunication(c *C) {
	request := CommunicationOperation{code: QUEUE_STATUS, channel: make(chan CommunicationOperation, 2)}
	s.communication <- request
	response := <- request.channel
	c.Assert(response.err, IsNil)
	c.Assert(response.code, Equals, QUEUE_STATUS_RESPONSE)
	c.Assert(response.message.(int), Equals, 5)
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


