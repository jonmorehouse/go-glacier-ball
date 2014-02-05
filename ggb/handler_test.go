package ggb

import (

	"testing"
	. "launchpad.net/gocheck"
)

// bootstrap handlers for this particular package component
func TestHandlers(t * testing.T) { 
	
	TestingT(t) 
}

// test the basic helper as needed
func (s * GGBSuite) TestHelper(c *C) {

	c.Check(42, Equals, 42)
}


