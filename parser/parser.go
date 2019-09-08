package parser

import (
	log "github.com/sirupsen/logrus"
)

var logMsg = log.Infof
var logErr = log.Errorf
var logWarn = log.Warnf
var logFatal = log.Fatalf

func Run(file string) {
	logMsg("Reading in file: %s", file)
}
