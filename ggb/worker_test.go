package ggb

import (
	"sync"
	. "launchpad.net/gocheck"
)

type WorkerSuite struct {

	comm chan CommunicationOperation
	fwg sync.WaitGroup
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
	s.comm = make(chan CommunicationOperation)
	s.files = CreateFileList(50)
	for i := range s.files {
		s.filePaths = append(s.filePaths, s.files[i].path)
	}
	go FileQueueManager(&s.fwg, s.comm)// master communication processor
	// now process them
	s.wg.Add(1)
	go ProcessorManager(&s.wg, &s.filePaths)// this queues up all the files
	s.wg.Wait()
}

func (s *WorkerSuite) TearDownTest(c *C) {

	RemoveFileList(&s.files)
	s.filePaths = []string{}//reset list - theres a more efficient way to do this 
}

func (s *WorkerSuite) TestWorker(c *C) {
	
	comm := make(chan CommunicationOperation)
	s.wg.Add(1)
	go Worker(&s.wg, comm)
	s.wg.Wait()
	// worker should be done processing all of the files
	// check the bucket for the tarball keys as needed
}

