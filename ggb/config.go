package ggb

import (
	"github.com/jonmorehouse/go-config/config"
)

func BootstrapConfig() {
	envVars := []string{
		"AWS_ACCESS_KEY_ID",
		"AWS_ACCESS_KEY_SECRET",
		"TARBALL_SIZE",
		"GO_PROCESSES",
	}
	config.New()//instantiate config package 
	config.Bootstrap(envVars) 
}

