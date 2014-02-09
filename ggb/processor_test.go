package ggb

import (

	"testing"
	"sync"
	. "launchpad.net/gocheck"
)

type ProcessorSuite struct {

	push chan PushOperation //where to push the files to 
	comm chan CommunicationOperation // used for managing errors taht are output
	waitGroup sync.WaitGroup
	files []*File // list of complete file structures etc
	filePaths []string // list of paths 

}

var _ = Suite(&ProcessorSuite{})

func (s *ProcessorSuite) SetUpSuite(c *C) {

	s.push = make(chan PushOperation, 100)
	s.comm = make(chan CommunicationOperation)

	// create the files as needed
	s.files = CreateFileList(5)

	// fill in filePaths as needed
	for i := range s.files {

		s.filePaths = append(s.filePaths, s.files[i].path)
	}

}

// 
func TestProcessor(t * testing.T) {

	TestingT(t)
}

func (s *ProcessorSuite) TestProcessor(c *C) {

	// test valid files as needed
	go Processor(&s.waitGroup, s.push, s.comm, s.filePaths)

	


}

func (s *ProcessorSuite) TestProcessorErrorHandling(c *C) {

	// delete all the files before running processor

}








