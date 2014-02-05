package ggb

import (
	
	. "launchpad.net/gocheck"
	"time"
	"testing"
	"code.google.com/p/go-uuid/uuid"
	"math/rand"
)

// create a custom suite for this element here
type FileSuite struct {

	parent * GGBSuite
	path string	
	size int // kilobytes
	large bool // whether or not the file should be flagged as large
}

// bootstrap file tests to use this suite as needed
func TestFiles(t * testing.T) {

	TestingT(t)
}

// now create the suite object
var _ = Suite(&FileSuite{})

// lets set up the file suite as needed
func (s * FileSuite) SetUpSuite(c *C) {

	// seed the random generator
	rand.Seed(time.Now().UTC().UnixNano())
	maxSize := 1024 * 1024

	// generate the file parameters
	s.path = uuid.New()
	s.size = 256 + rand.Intn(maxSize)

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

	


}

// 
func (s *FileSuite) TestFileSize(c *C) {

}


