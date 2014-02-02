package main

import (
	"fmt"
	"os"
	//"log"
	"io/ioutil"
)

func main() {

	bytes, err := ioutil.ReadAll(os.Stdin)
	
	//log.Println(err, string(bytes))
	if err != nil {

		fmt.Println(err)

	}

	fmt.Println(string(bytes))
}


