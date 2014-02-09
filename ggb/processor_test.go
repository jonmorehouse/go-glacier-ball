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

	var element PushOperation

	// test valid files as needed
	go Processor(&s.waitGroup, s.push, s.comm, s.filePaths)

	// wait for the waitgroup to finish
	s.waitGroup.Wait()
	
	// now lets loop through each of the paths
	for i := range s.filePaths {

		element = <- s.push

		// ensure that the files line up with the correct elements
		c.Assert(s.files[i].path, Equals, element.file.path)
		c.Assert(s.files[i].size, Equals, element.file.size)
	}
}

func (s *ProcessorSuite) TestProcessorErrorHandling(c *C) {

	var err error
	var commOperation CommunicationOperation

	// delete all the files before running processor
	for i := range s.filePaths {

		err = RemoveFile(s.filePaths[i])
		c.Assert(err, IsNil)
	}

	// test valid files as needed
	go Processor(&s.waitGroup, s.push, s.comm, s.filePaths)

	s.waitGroup.Wait()

	// now lets make sure we have errors as needed
	for _ = range s.files {

		commOperation = <- s.comm

		c.Assert(commOperation, Not(IsNil))
	}
}




