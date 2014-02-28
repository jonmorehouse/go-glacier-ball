package ggb

import (
	"io"
	. "launchpad.net/gocheck"
)

type UploaderSuite struct {
	GGBSuite
	reader io.Reader
}

var _ = Suite(&UploaderSuite{})

func (s *UploaderSuite) SetUpSuite(c *C) {
	
	s.Bootstrap()
	s.reader, _ = ReaderFromFilePaths(&s.filePaths)
}

func (s *UploaderSuite) TearDownSuite(c *C) {

	s.Breakdown()
}

func (s *UploaderSuite) TestUploader(c *C) {

	s.wg.Add(1)
	go FileUploader(&s.wg, s.reader)
	s.wg.Wait()
}


