package ggb

import (

	"os"
)

type File struct {

	path string // full path of the string as needed for opening / archiving
	size int64 // size in bytes

}

func NewFile (path string) (File, error) {

	// allocate a new file structure
	file := File{path: path}
	
	// grab the stat for the file
	stat, err := os.Stat(file.path)

	// now if there is an error return that
	if err != nil {

		return File{}, err
	}

	// grab the file size etc
	file.size = stat.Size()

	return file, nil
}


