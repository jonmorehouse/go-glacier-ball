package ggb

import (

	. "launchpad.net/gocheck"
	"testing"
	"fmt"

)

// bootstrap gocheck into all tests for this package as needed
func Test(t * testing.T) { TestingT(t) }

// generate teh test suite
type GGBSuite struct {
	
	// can allocate any variables needed here globally
	testDirectory string
}

// now we can have the ggb suit as our global suite
var _ = Suite(&GGBSuite{})

func (s *GGBSuite) TestHelloWorld(c *C) {
	
	c.Check(42, Equals, "42")
	fmt.Println("TEST")


}
