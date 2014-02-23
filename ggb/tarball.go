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

func NewTarball(prefix string, id int32) (*Tarball, error) {
	// generate keys and id for this file 
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

// Add a file and upload/create a new tarball if necessary 
func (t *Tarball) AddFile(file *File) (*Tarball, error) {
	// check to see if we have room in the current tarball
	if file.size > int64(config.Value("MAX_TARBALL_SIZE").(int)) || t.size + file.size > int64(config.Value("MAX_TARBALL_SIZE").(int)) {
		// create a single tarball upload for this element
		tarball, err := NewTarball(t.Prefix, atomic.AddInt32(&tarballCounter, 1))
		if err != nil {
			return nil, err
		} else {// add this larger file into the tarball as needed
			if err := tarball.addFile(file); err != nil {
				return nil, err
			} 
			if err := tarball.Upload(); err != nil {
				return nil, err
			}
			return tarball, nil
		}
	} else {// fits in our current archive. add it like normal
		if err := t.addFile(file); err != nil {
			return nil, err
		}
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

