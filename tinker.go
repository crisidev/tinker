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
	flagConfigFile = kingpin.Flag("config-file", "config file. mainly used for testing as config file is inside git-repo").Short('c').String()
	flagGitRepo    = kingpin.Flag("git-repo", "github repo for keys management").Required().Short('g').String()
	flagRunSetup   = kingpin.Flag("setup", "force recreation of private and public keys").Short('s').Bool()
)

func init() {
	LogSetup()
}

func main() {
	kingpin.Version(APP_VERSION)
	kingpin.Parse()
	config := TinkerConfiguration{}
	err := config.LoadConfig(*flagConfigFile)
	if err != nil {
		os.Exit(1)
	}
	if *flagRunSetup {
		if err = TincSetup(); err != nil {
			os.Exit(1)
		}
	}
	err = Cmd("tincd", "--help")
	os.Exit(0)
}
