package ggb

import (
	"os"
	"io"
	"strconv"
	"sync/atomic"
	"archive/tar"
	"compress/gzip"
	"github.com/jonmorehouse/go-config/config"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

const contentType = "application/xtar"
const permissions = s3.PublicReadWrite
const threshold = 0.25
var tarballCounter int32

type Tarball struct {
	// public components
	Id int32 // integer id of this particular element
	Key string // bucket key / local path while being created 
	Full bool // whether or not we still have room
	Prefix string // key prefix for this element
	// private components
	closed bool 
	size int64 // bytes
	file * os.File // currently open file
	gz * gzip.Writer // current gzip writer
	tw * tar.Writer // tarball writer that holds all of the tarball data as needed 
}

func NewTarball(prefix string) (*Tarball, error) {
	// increase the tarball counter as needed for this element
	id := atomic.AddInt32(&tarballCounter, 1)
	key := string(prefix) + strconv.Itoa(int(id)) + ".tar.gz" 
	file, err := os.Create(key)
	if err != nil {// return the error - our caller will pass to global error handler if necessary
		return nil, err
	}
	// create gzip compressor and link it to the file
	gz := gzip.NewWriter(file)
	// create the tarball writer and link it to gzip so all data is compressed 
	tw := tar.NewWriter(gz)
	// initialize the tarball with all of our variables as needed
	tarball := Tarball{
		Id: id, 
		Key: key,
		Prefix: prefix,
		Full: false,
		file: file,
		gz: gz,
		tw: tw,
	}
	return &tarball, nil
}

/*
	1.) check to see if new file will put us over 25% threshold
		yes: create a new tarball with only the new large file
		no: add file to the current tarball
	2.) guess which tarball of the two we should upload
	3.) if we are uploading the old tarball, then go ahead and return the new tarball

*/
func (t *Tarball) AddFile(file *File) (*Tarball, error) {
	
	var newTarball * Tarball
	maxSize := int64(config.Value("MAX_TARBALL_SIZE").(int))
	maxThresholdedSize := int64(float64(maxSize)*float64(1+threshold))
	// handle the adding of the file into the tarball
	// if file pushes current tarball over threshold or is big enough for its own tarball
	if file.size > maxSize || file.size + t.size > maxThresholdedSize {
		// create a new tarball
		tarball, err := NewTarball(t.Prefix)
		if err != nil {
			return nil, err	
		}
		if err := tarball.addFile(file); err != nil {
			return nil, err
		}
		newTarball = tarball 
	} else {//file size is small enough to add to the current tarball
		if err := t.addFile(file); err != nil {
			return nil, err
		}
	}
	// handle uploading / returning of the tarball
	if newTarball == nil {
		if t.size > maxSize {
			if err := t.Upload(); err != nil {
				return nil, err
			}
			return NewTarball(t.Prefix)
		} 
		return t, nil
	} else if t.size > maxSize && newTarball.size > maxSize { // upload both and return a new tarball (empty)
		if err := t.Upload(); err != nil {
			return nil, err
		}
		if err := newTarball.Upload(); err != nil {
			return nil, err
		}
		return NewTarball(t.Prefix)
	} else if t.size > newTarball.size { // upload t, return newTarball
		if err := t.Upload(); err != nil {
			return nil, err
		}
		return newTarball, nil
	} else if newTarball.size > t.size { // upload newTarball return t or nil
		if err := newTarball.Upload(); err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (t *Tarball) close() error {
	if t.closed {
		return nil
	}
	//close various file handles
	t.tw.Close()
	t.gz.Close()
	t.file.Close()
	// update status pointer	
	t.closed = true
	return nil
}

func (t *Tarball) addFile(file *File) error {
	// grab a handle on the file to read it into the buffer
	fr, err := os.Open(file.path)
	if err != nil {
		return err
	}
	defer fr.Close()
	if stat, err := fr.Stat(); err == nil {
		// create the tarball header for this file
		header := new(tar.Header)
		header.Name = file.path
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()
		// write the header to the current tarball
		if err := t.tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data over now that the header is set
		if _, err := io.Copy(t.tw, fr); err == nil {
			t.size += stat.Size()
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}

func (t *Tarball) Delete() error {
	
	if err := os.Remove(t.Key); err != nil {
		return nil
	}
	return nil
	
}

func (t *Tarball) Upload() error {
	// make sure the tarball is completed etc
	t.close()
	file, err := os.Open(t.Key)
	if err != nil {
		return err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	// now set up the bucket and prepare for the upload
	s3Conn := s3.New(config.Value("AWS_AUTH").(aws.Auth), config.Value("AWS_REGION").(aws.Region))
	bucket := s3Conn.Bucket(config.Value("BUCKET_NAME").(string))
	// now lets upload the file
	if err := bucket.PutReader(t.Key, file, stat.Size(), contentType, permissions); err != nil {
		return err
	}
	return nil
}

