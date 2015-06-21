package main

import (
	"log/syslog"
)

type SystemLogger interface {

	Emerg(string) error
	Crit(string) error
	Alert(string) error
	Err(string) error
	Warning(string) error
	Notice(string) error
	Info(string) error
	Debug(string) error
}

func CreateLogger(out string) (SystemLogger, error) {
	
	/*
	switch out {

	case "stdio":
		return log.New()
	
	case "syslog":

	}*/

	return syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, GetProgramName())
	
	//return nil, errors.New("Unknown logging implementation: " + out)
}
