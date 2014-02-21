package ggb

import (
	. "launchpad.net/gocheck"
	"strconv"
	"github.com/jonmorehouse/go-config/config"
)

type TarballSuite struct {
	files []*File
	filePaths []string 
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
	RemoveFiles(&s.files)
}

func (s *TarballSuite) TestNewTarball(c *C) {
	id := int32(3285)
	tarball, err := NewTarball(config.Value("TARBALL_PREFIX").(string), id)
	c.Assert(tarball, NotNil)
	c.Assert(err, IsNil)
	// now lets make sure that we have set the parameters properly
	c.Assert(tarball.Id, Equals, id)
	c.Assert(tarball.Key, Equals, config.Value("TARBALL_PREFIX").(string) +  strconv.Itoa(int(id)) + ".tar.gz")
	c.Assert(tarball.Full, Equals, false)
	c.Assert(tarball.file, NotNil)
	c.Assert(tarball.gz, NotNil)
	c.Assert(tarball.tw, NotNil)
}

func (s *TarballSuite) TestUpload(c *C) {


}

func (s *TarballSuite) TestAddFile(c *C) {


}

func (s *TarballSuite) TestaddFile(c *C) {


}


