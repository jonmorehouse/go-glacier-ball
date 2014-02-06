package ggb

import (
	"math/rand"
	. "launchpad.net/gocheck"
	"os"
	"time"
)

// generate teh test suite
type GGBSuite struct {
	
	// can allocate any variables needed here globally
	testDirectory string
}

// now we can have the ggb suit as our global suite
//var _ = Suite(&GGBSuite{})

func (s * GGBSuite) SetUpSuite(c *C) {

	// seed the random generator
	rand.Seed(time.Now().UTC().UnixNano())
}

func (s * GGBSuite) TearDownSuite(c *C) {


}

// generate any global helper functions on this element as needed etc
// create a dudd file with a bunch of random bytes
// size in kilobytes
func CreateFile(path string, size int64) error {

	var offset int64
	offset = 0

	// create the file at the path
	file, err := os.Create(path)

	// check if the err exists
	if err != nil {

		return err
	}

	// initialize kilobyte elements
	byteArray := make([]byte,1)

	// now lets loop through and write a byte to the file as needed
	for i := 0; int64(i) < size; i++ {

		_, err := file.WriteAt(byteArray, offset)
		offset += 1

		if err != nil {

			return err
		}
	}

	return nil
}

func RemoveFile(path string) error {

	// remove file
	err := os.Remove(path)

	if err != nil {

		return err

	}

	return nil
}

