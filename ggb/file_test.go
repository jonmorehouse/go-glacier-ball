package ggb

import (
	
	. "launchpad.net/gocheck"
	"code.google.com/p/go-uuid/uuid"
	"math/rand"
)

// create a custom suite for this element here
type FileSuite struct {

	path string	
	size int64 // size in bytes
	large bool // whether or not the file should be flagged as large
}

// lets set up the file suite as needed
func (s * FileSuite) SetUpSuite(c *C) {

	var maxSize int64
	maxSize = 512 // maximum number of bytes

	// generate the file parameters
	s.path = uuid.New() + ".txt"
	s.size = int64(256) + rand.Int63n(maxSize)

	// ensure that for test purposes, we have flagged the file handle properly
	if s.size < maxSize {

		s.large = false

	} else {

		s.large = true
	}

	// now lets go ahead and create the file as needed
	// now lets ensure that we don't have any errors being passed back to the elements calling as needed
	err:= CreateFile(s.path, s.size)

	// make sure that we don't have an error here from the creation of the file
	c.Assert(err, IsNil)
}

func (s * FileSuite) TearDownSuite(c *C) {

	err := RemoveFile(s.path)	
	c.Assert(err, IsNil)
}

// create the dudd file - and then create a file struct - and then compare size etc  
func (s *FileSuite) TestFileSize(c *C) {

	// now lets create a file and ensure that its the correct size
	file, err := NewFile(s.path)

	// make sure that we didn't get an error in return
	c.Assert(err, IsNil)	

	// now lets make sure that the file path is correct
	c.Assert(file.size, Equals, s.size)	
}


