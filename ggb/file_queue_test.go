package ggb

import (

	"code.google.com/p/go-uuid/uuid"
	"testing"
	"fmt"
	. "launchpad.net/gocheck"
)

// bootstrap gocheck into current test environment
func TestFileQueue(t * testing.T) {

	TestingT(t)
}

// now lets create a few different files 
type FileQueueSuite struct {

	files[]* File
}

// bootstrap the suite into the current environment
var _ = Suite(&FileQueueSuite{})

func (s * FileQueueSuite) SetUpSuite(c *C) {

	multiple := 256 * 1024 // kilobytes

	// create a list of files
	for i := 0; i < 5; i++ {

		// create a file of
		path := uuid.New() + ".txt"

		// initialize size of this file
		size := int64(multiple * i)

		// create the sample file as needed
		err := CreateFile(path, size)

		fmt.Println(err)

		// make sure no error created with files
		//c.Assert(err, Equals, IsNil)

		// now that our file is created, createa  File struct to contain it
		// we're storing pointers to these 
		fmt.Println(size)

	}
}

func (s *FileQueueSuite) TestFileQueue(c *C) {




}



