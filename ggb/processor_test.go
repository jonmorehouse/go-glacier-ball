package ggb

import (
	"sync"
	. "launchpad.net/gocheck"
	"io"
)

type ProcessorSuite struct {

	GGBSuite
	reader io.Reader 
	inputChannel chan string
	wg sync.WaitGroup
}

var _ = Suite(&ProcessorSuite{})

func (s *ProcessorSuite) SetUpTest(c *C) {

	s.Bootstrap()
	s.reader, _ = ReaderFromFilePaths(&s.filePaths)
	s.inputChannel = make(chan string)
	go FileQueueManager(&sync.WaitGroup{}, errorComm)// this is really irrelevant for this test
}

func (s *ProcessorSuite) TearDownTest(c *C) {

	s.Breakdown()
}

// add a bunch of files and ensure that they make it to the file queue 
func (s *ProcessorSuite) TestProcessor(c *C) {

	length := 0
	worker := func(wg * sync.WaitGroup, file *File) {

		defer wg.Done()
		c.Assert(file, NotNil)
		c.Assert(file.path, NotNil)
		length += 1
	}

	// call the process as needed
	s.wg.Add(1)
	go ProcessorManager(&s.wg, s.reader)
	s.wg.Wait()
	// now lets ensure all the files got queued as needed
	ProcessFileQueue(worker)
	// make sure all files got queued as needed
	c.Assert(length, Equals, len(s.filePaths))

}


