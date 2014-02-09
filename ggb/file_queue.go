package ggb


func NewFileQueueItem(file * File) FileQueueItem {

	// create a new item and set priority to the file size
	item := FileQueueItem{
		file: file,
		priority: file.size,
	}

	return item
}





