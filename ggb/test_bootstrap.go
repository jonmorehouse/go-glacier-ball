package ggb

import (
	"sync"
	"code.google.com/p/go-uuid/uuid"
	"math/rand"
	"os"
	"bytes"
)

type GGBSuite struct {

	comm chan CommunicationOperation
	wg sync.WaitGroup
	numberFiles int
	files []*File
	filePaths []string
}

func (s *GGBSuite) Bootstrap() {

	Bootstrap()//call master bootstrap method
	if s.numberFiles == 0 {
		s.numberFiles = 5
	} else if s.numberFiles < 0 {
		return
	}
	s.files = CreateFileList(s.numberFiles)
	for i := range s.files {
		s.filePaths = append(s.filePaths, s.files[i].path)
	}
}

func (s *GGBSuite) Breakdown() {

	RemoveFileList(&s.files)
	s.filePaths = []string{}
}


func CreateFileList(quantity int) []*File {
	var path string
	var size int64
	files := make([]*File, quantity)
	for i := 0; i < quantity; i++ {
		path = uuid.New() + ".txt"
		size = int64(256) + rand.Int63n(1024)
		CreateFile(path, size)
		file, _ := NewFile(path)
		files[i] = &file
	}
	return files
}


// generate any global helper functions on this element as needed etc
// create a dudd file with a bunch of random bytes
func CreateFile(path string, size int64) error {
	file, err := os.Create(path)
	// check if the err exists
	if err != nil {
		return err
	}
	byteArray := make([]byte,1)
	_, err = file.WriteAt(byteArray, size)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(path string) error {

	// remove file
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFileList(files * []*File) {
	for i := range *files {
		_ = RemoveFile((*files)[i].path)	
	}
	*files = []*File{}
}

func ReaderFromFilePaths(filePaths * []string) (* bytes.Reader, error) {
	
	buffer := bytes.NewBuffer([]byte{})
	// create a reader with files
	for i := range (*filePaths) {
		buffer.WriteString((*filePaths)[i])
		buffer.WriteString("\n")
	}
	// return a new reader with this string
	return bytes.NewReader(buffer.Bytes()), nil
}

func ReaderFromString(input string) (*bytes.Reader, error) {

	buffer := bytes.NewBuffer([]byte{})
	_, err := buffer.WriteString(input)
	if err != nil {

		return nil, err

	}
	return bytes.NewReader(buffer.Bytes()), nil
}

