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

/*

	1.) Split files into 5 slices - pass to processors
	2.) Start workers (they have their own waitgroup) 
	3.) Waitgroup waits on the processors to complete
	4.) Waitgroup waits on the fileQueue to stop 
	5.) Once we know the fileQueue is finalized, pass a message to all of our worker queues - tell them its the final run!
	6.) Wait on all workers to finish!
*/



