package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	APP_NAME    = "tinker"
	APP_VERSION = "0.1"
	APP_SITE    = "https://github.com/crisidev/tinker"
)

var (
	flagDebug      = kingpin.Flag("debug", "enable debug mode").Short('d').Bool()
	flagLogFile    = kingpin.Flag("log-file", "enable logging to file").Short('l').Default("").String()
	flagConfigFile = kingpin.Flag("config-file", "enable logging to file").Short('c').Default("config.json").String()
)

func init() {
	LogSetup()
}

func main() {
	kingpin.Version(APP_VERSION)
	kingpin.Parse()
	config := TinkerConfiguration{}
	err := config.LoadFromFile(*flagConfigFile)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}