package ggb

import (

	"testing"
	"sync"
	. "launchpad.net/gocheck"
)

type FileQueueSuite struct {

	push chan PushOperation
	pop chan PopOperation
	communication chan CommunicationOperation
	waitGroup sync.WaitGroup

}

// bootstrap the suite into the current environment
var _ = Suite(&FileQueueSuite{})

// bootstrap gocheck into current test environment
func Test(t * testing.T) {

	TestingT(t)
}

// now lets test the basic functionality of the different elements as needed
func (s * FileQueueSuite) SetUpSuite(c *C) {

	s.push = make(chan PushOperation)
	s.pop = make(chan PopOperation)
	s.communication = make(chan CommunicationOperation)

	// now we can start the go routine as needed
	go FileQueueManager(&s.waitGroup, s.push, s.pop, s.communication)

}

func (s *FileQueueSuite) TestSuccessfulQueue(c *C) {

	// initialize channel operations needed  

	// push in a dudd file
	push := PushOperation{file: &File{path: "test.txt", size: 10000}}

	// initialize a pop operation - this should allow us to grab a file
	pop := PopOperation{channel: make(chan *File, 1)}

	// communication operation - signify end of submissions 
	comm := CommunicationOperation{code: ALL_JOBS_SUBMITTED}

	// now lets check the status element as needed
	status := CommunicationOperation{code: QUEUE_STATUS, channel: make(chan CommunicationOperation, 1)}

	// pump the operations into our worker
	s.push <- push
	s.communication <- status
	s.pop <- pop
	s.communication <- comm

	// wait until the queueManager is finished 
	s.waitGroup.Wait()

	// now that we are finished - lets assert that the popChannel has something buffered 
	fileResponse := <- pop.channel
	statusResponse := <- status.channel 

	// now lets make sure that the file is not nil 
	c.Assert(fileResponse, Not(Equals), nil)
	c.Assert(fileResponse.path, Not(Equals), nil)

	// check out status response
	c.Assert(statusResponse, Not(Equals), nil)
	
	// now ensure that we got the correct response for the status back
	c.Assert(statusResponse.message.(int), Equals, 1)

}


