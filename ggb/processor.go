package ggb

import (
	"sync"
	//"github.com/jonmorehouse/go-config/config"
	"bufio"
	"io"
	"strings"
)

func ProcessorManager(cwg * sync.WaitGroup, rawInput io.Reader) {

	defer cwg.Done()
	var wg sync.WaitGroup
	reader := bufio.NewReader(rawInput)
	channel := make(chan string)
	// start process worker
	wg.Add(1)
	go Processor(&wg, channel)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.Trim(line, "\n")
		channel <- line
	}
	// now close the channel as needed
	close(channel)
	// now wait until all files are closed and submitted (in the processor)
	wg.Wait()
}

func Processor(cwg * sync.WaitGroup, input chan string) {

	var wg sync.WaitGroup
	for filePath := range input {
		wg.Add(1)
		go func() {
			// create the file as needed
			file, err := NewFile(filePath)
			if err != nil {
				// pass the error to our error handler
				errorComm <- CommunicationOperation{code: ERROR_FILE, message: &err}
			} else {
				// push this file to the push channel as needed
				push <- PushOperation{file: &file}
			}
			wg.Done()
		}()
	}
	// make sure all processes finish 
	wg.Wait()
	cwg.Done()
}

