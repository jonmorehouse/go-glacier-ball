package ggb

import (

	//"container/heap"
)

type FileQueueItem struct {

	file * File // file   	

	priority int64 // this is going to be based upon the 

	index int

}

func NewFileQueueItem(file * File) FileQueueItem {

	// create a new item and set priority to the file size
	item := FileQueueItem{
		file: file,
		priority: file.size,
	}

	return item
}





