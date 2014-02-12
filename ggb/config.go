package ggb

// global variables needed throughout application
type Config struct {

	// aws access credentials
	AWS_ACCESS_KEY_ID string
	AWS_ACCESS_KEY_SECRETE string

	// max tarball size
	TARBALL_SIZE int64
}

// now lets create the config element as needed
var config * Config = nil


