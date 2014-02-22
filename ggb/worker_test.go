package ggb

import (
	//"sync"
	. "launchpad.net/gocheck"
)

type WorkerSuite struct {

	files []*File
}

var _ = Suite(&WorkerSuite{})

func (s *WorkerSuite) SetUpSuite(c *C) {

	Bootstrap()

}

func (s *WorkerSuite) SetUpTest(c *C) {

	s.files = CreateFileList(50)

}

func (s *WorkerSuite) TearDownSuite(c *C) {

}

func (s *WorkerSuite) TearDownTest(c *C) {


}

func (s *WorkerSuite) TestWorker(c *C) {


}



