package ggb

import (
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

	Bootstrap()
}

func (s *ProcessorSuite) SetUpTest(c *C) {

	s.comm = make(chan CommunicationOperation)
	push = make(chan PushOperation, 1000)

	// create the files as needed
	s.files = CreateFileList(5)

	// fill in filePaths as needed
	for i := range s.files {

		s.filePaths = append(s.filePaths, s.files[i].path)
	}
}

func (s *ProcessorSuite) TearDownTest(c *C) {

	for i := range s.filePaths {
		_ = RemoveFile(s.filePaths[i])
	}

	s.files = []*File{}
	s.filePaths = []string{}
}


func (s *ProcessorSuite) TestProcessor(c *C) {
	var element PushOperation
	// test valid files as needed
	go Processor(&s.waitGroup, s.comm, s.filePaths)
	// wait for the waitgroup to finish
	s.waitGroup.Wait()
	// now lets loop through each of the paths
	for i := range s.filePaths {
		element = <- push
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
	go Processor(&s.waitGroup, s.comm, s.filePaths)
	s.waitGroup.Wait()
	// now lets make sure we have errors as needed
	for _ = range s.files {
		commOperation = <- s.comm
		c.Assert(commOperation, Not(IsNil))
	}
}


func (s *ProcessorSuite) TestProcessorManager(c *C) {
	filePaths := []string{}
	// copy old paths into the new list 
	for i := 0; i < 50; i++ {
		filePaths = append(filePaths, s.filePaths...)
	}
	s.filePaths = make([]string, len(filePaths))
	copy(s.filePaths, filePaths)
	// processor manager is responsible for booting up and managing all file process workers 
	go ProcessorManager(&s.waitGroup, &filePaths)
	// wait for waitgroup to finish before running any tests
	s.waitGroup.Wait()
	// verify that all files needed were properly queued up etc
	for i := 0; i < len(s.filePaths); i++ /*range push*/ {
		input := <- push
		c.Assert(input.file, Not(Equals), nil)
	}
}


