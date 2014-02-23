package ggb

import (
	"fmt"
	"sync"
	. "launchpad.net/gocheck"
)

type WorkerSuite struct {

	wg sync.WaitGroup
	files []*File
	filePaths []string
}

var _ = Suite(&WorkerSuite{})

func (s *WorkerSuite) SetUpSuite(c *C) {

	Bootstrap()
}

func (s *WorkerSuite) SetUpTest(c *C) {
	// generate the files 
	s.files = CreateFileList(50)
	for i := range s.files {
		s.filePaths = append(s.filePaths, s.files[i].path)
	}
	// now process them
	go ProcessorManager(&s.wg, &s.filePaths)
	s.wg.Wait()
}

func (s *WorkerSuite) TearDownTest(c *C) {

	RemoveFileList(&s.files)
	s.filePaths = []string{}//reset list - theres a more efficient way to do this 
}

func (s *WorkerSuite) TestWorker(c *C) {

	popResponseChannel := make(chan PopResponseOperation, 10)
	c.Assert(popResponseChannel, NotNil)

	// submit a pop request
	pop <- PopOperation{channel: popResponseChannel}
	popResponse := <- popResponseChannel

	fmt.Println(popResponse)


}



