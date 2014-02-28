package ggb

import (
	"io"
	. "launchpad.net/gocheck"
)

type UploaderSuite struct {
	GGBSuite
	input io.Reader
}

var _ = Suite(&UploaderSuite{})

func (s *UploaderSuite) SetUpSuite(c *C) {
	
	s.Bootstrap()
	// create input as needed
	reader, err := ReaderFromFilePaths(&s.filePaths)
	c.Assert(err, IsNil)
	c.Assert(reader, NotNil)
	s.input = reader
	
}

func (s *UploaderSuite) TearDownSuite(c *C) {

	s.Breakdown()

}

func (s *UploaderSuite) TestUploader(c *C) {

	s.wg.Add(1)
	go FileUploader(&s.wg, s.input)
	s.wg.Wait()

}


