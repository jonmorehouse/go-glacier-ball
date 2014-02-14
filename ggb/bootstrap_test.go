package ggb

import (
	. "launchpad.net/gocheck"
	"testing"
	"fmt"
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
	errorOperation := CommunicationOperation{}

}


