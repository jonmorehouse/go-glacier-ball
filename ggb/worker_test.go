package ggb

import (
	"sync"
	. "launchpad.net/gocheck"
)

type WorkerSuite struct {

	GGBSuite
	wg sync.WaitGroup
}

var _ = Suite(&WorkerSuite{})

func (s *WorkerSuite) SetUpSuite(c *C) {

	s.Bootstrap()
	reader,_ := ReaderFromFilePaths(&s.filePaths)
	// initialize queue and process files 
	go FileQueueManager(&sync.WaitGroup{}, errorComm)
	s.wg.Add(1)
	go ProcessorManager(&s.wg, reader)
	s.wg.Wait()
}

func (s *WorkerSuite) TearDownSuite(c *C) {

}

func (s *WorkerSuite) TestWorker(c *C) {


}


