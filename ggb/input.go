package ggb

import (
	"bufio"
	"io"
	"os"
	"errors"
)

const (
	INPUT_TYPE_NULL = 00
	INPUT_TYPE_STRING = 01
	INPUT_TYPE_FILE = 02 
	INPUT_TYPE_UNKNOWN = 03
)

/*
	file rules
	a single line that starts with a dot and no space
	lines with a .2char or .3char at the end of the line

	loop through the first few lines and see if we have 
*/
func InputType(ioReader io.ReadSeeker) (int, error) {

	reader := bufio.NewReader(ioReader)
	line, err := reader.ReadString('\n')	
	defer ioReader.Seek(int64(0), 0)
	if err != nil && len(line) == 0 {//eof -- nothing passed in
		return INPUT_TYPE_NULL, errors.New("No input")
	}
	// now we want to check if the file exists
	if _, err := os.Stat(line); err == nil {//file exists
		return INPUT_TYPE_FILE, nil
	} else {
		return INPUT_TYPE_STRING, nil
	}
	return INPUT_TYPE_UNKNOWN, errors.New("Couldn't determine input type")
}





