package ggb

import (
	. "launchpad.net/gocheck"
	"testing"
)

type BootstrapSuite struct {}

var _ = Suite(&BootstrapSuite{})
func TestBootstrap(t * testing.T) {
	TestingT(t)
}

func (s *BootstrapSuite) SetUpSuite(c *C) {
	Bootstrap()
}

func (s *BootstrapSuite) TestErrorHandler(c *C) {
	// create a nonsimple error operation
	errorOperation := CommunicationOperation{}
	// start go worker for error handling - need to figure out a nice way to handle tests on this element as needed
	go ErrorHandler()
	errorComm <- errorOperation
}


