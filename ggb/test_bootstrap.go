package ggb

import (

	. "launchpad.net/gocheck"
	"os"
)

// generate teh test suite
type GGBSuite struct {
	
	// can allocate any variables needed here globally
	testDirectory string

}

// now we can have the ggb suit as our global suite
var _ = Suite(&GGBSuite{})

func (s * GGBSuite) SetUpSuite(c *C) {


}

func (s * GGBSuite) TearDownSuite(c *C) {



}

// generate any global helper functions on this element as needed etc

// create a dudd file with a bunch of random bytes
// size in kilobytes
func CreateFile(path string, size int) error {

	// now lets loop through the size and write empty bytes to the element
	file, err := os.Open(path)

	// catch and delegate error here as needed
	if err != nil {

		return err
	}

	// now lets loop through and write a byte to the file as needed

	// now lets go 
	return nil
}

func RemoveFile(path string) {

	// 

}





