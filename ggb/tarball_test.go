package ggb

import (

	//"sync"
	. "launchpad.net/gocheck"

)

type TarballSuite struct {
	files []*File
}

var _ = Suite(&TarballSuite{})

func (s *TarballSuite) SetUpSuite(c *C) {

	s.files = CreateFileList(50)
}

func (s *TarballSuite) TearDownSuite(c *C) {

}



