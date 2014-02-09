package ggb

import (

	"container/list"
	"code.google.com/p/go-uuid/uuid"
	"testing"
	. "launchpad.net/gocheck"
)

// bootstrap the suite into the current environment
var _ = Suite(&FileQueueSuite{})

// bootstrap gocheck into current test environment
func Test(t * testing.T) {

	TestingT(t)
}

// now lets create a few different files 
type FileQueueSuite struct {

	files list.List
}

func (s * FileQueueSuite) SetUpSuite(c *C) {

	multiple := 256 * 1024 // kilobytes

	// create a list of files
	for i := 1; i <= 5; i++ {

		// create a file of
		path := "/tmp/go/" + uuid.New() + ".txt"

		// initialize size of this file
		size := int64(multiple * i)

		// create the sample file as needed
		err := CreateFile(path, size)

		// make sure no error created with files
		c.Assert(err, IsNil)

		// now that our file is created, createa  File struct to contain it
		// we're storing pointers to these 
		file,err := NewFile(path)

		c.Assert(err, IsNil)

		s.files.PushBack(&file)
	}
}

func (s * FileQueueSuite) TearDownSuite(c *C) {

	// remove all file objects as needed
	for e := s.files.Front(); e != nil; e = e.Next() {

		s.files.Remove(e)

	}
}

func (s *FileQueueSuite) TestFileQueue(c *C) {

	for e := s.files.Front(); e != nil; e = e.Next() {

		// now lets queue these files up
		
		
	}
}





