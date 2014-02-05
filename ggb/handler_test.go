package ggb

import (

	. "launchpad.net/gocheck"
	"fmt"
)

func (s * GGBSuite) TestTest(c *C) {

	fmt.Println("TEST")
	c.Check(42, Equals, 40)
}


