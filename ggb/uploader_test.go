package ggb

import (
	"fmt"
	. "launchpad.net/gocheck"
)

type UploaderSuite struct {
	GGBSuite
}

var _ = Suite(&UploaderSuite{})

func (s *UploaderSuite) SetUpSuite(c *C) {

	s.Bootstrap()
}

func (s *UploaderSuite) TearDownSuite(c *C) {

	s.Breakdown()

}

func (s *UploaderSuite) TestUploader(c *C) {

	// 
	fmt.Println("ASDF")
}




