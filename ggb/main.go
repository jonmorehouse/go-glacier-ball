package ggb

import (
	"os"
	"bufio"
	"log"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		line, err := reader.ReadString('\n')
		
		if err != nil {

			break

		}

		log.Println(line)
	}
	
}




