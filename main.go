package main

import (
	"fmt"
	"labix.org/v2/pipe"
)

func main() {

	p := pipe.Line(
	
		pipe.ReadFile("test.txt"),
		pipe.Exec("lpr"),

	)

	output, err := pipe.CombinedOutput(p)

	if err != nil {

		fmt.Printf("%v\n", err)

	}

	fmt.Printf("%s", output)
}


