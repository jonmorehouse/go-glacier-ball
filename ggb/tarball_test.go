package ggb

import (
	. "launchpad.net/gocheck"
	"strconv"
	"github.com/jonmorehouse/go-config/config"
)

type TarballSuite struct {
	files []*File
	filePaths []string 
	tarball * Tarball
}

var _ = Suite(&TarballSuite{})

func (s *TarballSuite) SetUpSuite(c *C) {
	Bootstrap()
	s.files = CreateFileList(50)
	for i := range s.files {
		s.filePaths = append(s.filePaths, s.files[i].path)
	}
}

func (s *TarballSuite) TearDownSuite(c *C) {
	RemoveFileList(&s.files)
	s.filePaths = []string{}
}

func (s *TarballSuite) TearDownTest(c *C) {
	if s.tarball != nil {
		err := s.tarball.Delete()
		c.Assert(err, IsNil)
	}
}

func (s *TarballSuite) TestNewTarball(c *C) {
	tarball, err := NewTarball(config.Value("TARBALL_PREFIX").(string))
	s.tarball = tarball
	c.Assert(tarball, NotNil)
	c.Assert(err, IsNil)
	// now lets make sure that we have set the parameters properly
	c.Assert(tarball.Id, Equals, tarballCounter)
	c.Assert(tarball.Key, Equals, config.Value("TARBALL_PREFIX").(string) +  strconv.Itoa(int(tarballCounter)) + ".tar.gz")
	c.Assert(tarball.Full, Equals, false)
	c.Assert(tarball.file, NotNil)
	c.Assert(tarball.gz, NotNil)
	c.Assert(tarball.tw, NotNil)
}

func (s *TarballSuite) TestUpload(c *C) {

	tarball, err := NewTarball(config.Value("TARBALL_PREFIX").(string))
	s.tarball = tarball
	c.Assert(err, IsNil)
	c.Assert(tarball, NotNil)
	// add a file to the archive
	newTarball, err := tarball.AddFile(s.files[0])
	c.Assert(newTarball, IsNil)
	c.Assert(err, IsNil)
	// now we can upload the archive
	err = tarball.Upload()
	c.Assert(err, IsNil)
}


