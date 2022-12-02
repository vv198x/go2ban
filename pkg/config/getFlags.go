package config

import (
	"flag"
)

const (
	defaultTmp = "/tmp"
	defaultLog = "./"
)

var (
	tmpDir = flag.String("tmpDir", defaultTmp, "Directory for cache scanners")
	logDir = flag.String("logDir", defaultLog, "Directory for logs")
)

func Get(flag string) string {
	switch flag {
	case "tmpDir":
		return *tmpDir
	case "logDir":
		return *logDir
	}
	return ""
}
