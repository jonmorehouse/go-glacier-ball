package ggb

import (
	. "launchpad.net/gocheck"
	"bytes"
	//"fmt"
)

type InputSuite struct {
	
	GGBSuite	
}

var _ = Suite(&InputSuite{})

func (s *InputSuite) SetUpSuite(c *C) {
	s.Bootstrap()
}

func (s *InputSuite) TearDownSuite(c *C) {

	s.Breakdown()
}

func (s *InputSuite) TestFileInput(c *C) {

	// initialize the reader with our file paths as needed
	reader, err := ReaderFromFilePaths(&s.filePaths)
	c.Assert(err, IsNil)
	c.Assert(reader, NotNil)
}

func (s *InputSuite) TestTextInput(c *C) {

	// initialize string text to test inputType
	stringContents := "hello world this is a text message"
	reader, err := ReaderFromString(stringContents)
	c.Assert(err, IsNil)
	c.Assert(reader, NotNil)
	// test input type to be string
	inputType, err := InputType(reader)
	c.Assert(err, IsNil)
	c.Assert(inputType, Equals, INPUT_TYPE_STRING)

	// make sure test input resets the reader
	output := make([]byte, len(stringContents))
	reader.Read(output)
	c.Assert(string(output), Equals, stringContents)
}

func (s *InputSuite) TestNullInput(c *C) {
	// create an empty buffer
	buffer := bytes.NewBuffer([]byte{})
	reader := bytes.NewReader(buffer.Bytes())
	c.Assert(reader, NotNil)
	// now check input type results
	inputType, err := InputType(reader)
	c.Assert(err, NotNil)
	c.Assert(inputType, Equals, INPUT_TYPE_NULL)
}





