package ggb

import (
	"fmt"
	"sync"
	. "launchpad.net/gocheck"
)

type UploaderSuite struct {
	
	wg sync.WaitGroup
	files []*File
	filePaths []string
	
}

var _ = Suite(&UploaderSuite{})

func (s *UploaderSuite) SetUpSuite(c *C) {
	Bootstrap()
}

func (s *UploaderSuite) TestTest(c *C) {

	fmt.Println("ASDFASDF")


}





