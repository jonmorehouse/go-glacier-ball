package ggb

import (

	"code.google.com/p/go-uuid/uuid"
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

func CreateFileList(quantity int) []*File {

	var path string
	var size int64

	// create files array
	files := make([]*File, quantity)

	// loop through and create all files as needed
	for i := 0; i < quantity; i++ {

		path = uuid.New() + ".txt"
		size = int64(256) + rand.Int63n(1024)

		// create file on disk
		CreateFile(path, size)

		// create file object
		file, _ := NewFile(path)

		// now that we have the file path created
		files[i] = &file
	}

	return files
}


// generate any global helper functions on this element as needed etc
// create a dudd file with a bunch of random bytes
func CreateFile(path string, size int64) error {

	// create the file at the path
	file, err := os.Create(path)

	// check if the err exists
	if err != nil {

		return err
	}

	// initialize kilobyte elements
	byteArray := make([]byte,1)

	_, err = file.WriteAt(byteArray, size)

	if err != nil {
	
		return err
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

