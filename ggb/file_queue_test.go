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
	//go FileQueue(&s.waitGroup, s.push, s.pop, s.communication)

}

func (s *FileQueueSuite) TestSuccessfulQueue(c *C) {

	push := PushOperation{message: "test push message"}
	pop := PopOperation{message: "test pop message"}
	comm := CommunicationOperation{code: ALL_JOBS_SUBMITTED}

	s.push <- push
	s.pop <- pop
	s.communication <- comm

	s.waitGroup.Wait()

}




