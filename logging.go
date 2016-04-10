package main

import (
	"os"

	"github.com/op/go-logging"
)

var (
	log           = logging.MustGetLogger(APP_NAME)
	logFormat     = logging.MustStringFormatter(`%{color}%{time:15:04:05} %{shortfile:.24s} %{shortfunc:.16s} %{level:.5s} %{id:03x}%{color:reset} %{message}`)
	logFileFormat = logging.MustStringFormatter(`%{time:15:04:05} %{shortfile:.24s} %{shortfunc:.16s} %{level:.5s} %{id:03x} %{message}`)
)

// Setup logging using go-logging (see https://github.com/op/go-logging)
// Global variables flagDebug and flagLogFile are used to control logging configuration:
// - default logging is on Stderr and logging level is WARN
// - flagLogFile set to a valid path setup a file logger with minumum level INFO
// - flagDebug increases default and logFile logger level to DEBUG
// There is not defer file.Close() because it breaks go-logging and the logfile needs to
// be closed at the end of the program.
func LogSetup() (err error) {
	outBackend := logging.NewLogBackend(os.Stderr, "", 0)
	outBackendFormatter := logging.NewBackendFormatter(outBackend, logFormat)
	outBackendLeveled := logging.AddModuleLevel(outBackendFormatter)
	if *flagDebug {
		outBackendLeveled.SetLevel(logging.DEBUG, "")
	} else {
		outBackendLeveled.SetLevel(logging.INFO, "")
	}

	if *flagLogFile != "" {
		file, err := os.OpenFile(*flagLogFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Errorf("error opening logfile %s: %v", *flagLogFile, err)
			return err
		}
		fileBackend := logging.NewLogBackend(file, "", 0)
		fileBackendFormatter := logging.NewBackendFormatter(fileBackend, logFileFormat)
		fileBackendLeveled := logging.AddModuleLevel(fileBackendFormatter)
		if *flagDebug {
			fileBackendLeveled.SetLevel(logging.DEBUG, "")
		} else {
			fileBackendLeveled.SetLevel(logging.INFO, "")
		}
		logging.SetBackend(outBackendLeveled, fileBackendLeveled)
	} else {
		logging.SetBackend(outBackendLeveled)
	}
	return nil
}
