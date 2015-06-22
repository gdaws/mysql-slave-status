package main

import (
	"errors"
	"fmt"
	"io"
	"log/syslog"
	"os"
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

type WriterSystemLogger struct {
	writer io.Writer
}

func NewWriterSystemLogger(writer io.Writer) *WriterSystemLogger {

	return &WriterSystemLogger{
		writer: writer,
	}
}

func (logger *WriterSystemLogger) writeMessage(priority string, message string) error {
	_, err := logger.writer.Write([]byte(fmt.Sprintln(priority + ": " + message)))
	return err
}

func (logger *WriterSystemLogger) Emerg(message string) error {
	return logger.writeMessage("EMERG", message)
}

func (logger *WriterSystemLogger) Crit(message string) error {
	return logger.writeMessage("CRIT", message)
}

func (logger *WriterSystemLogger) Alert(message string) error {
	return logger.writeMessage("ALERT", message)
}

func (logger *WriterSystemLogger) Err(message string) error {
	return logger.writeMessage("ERR", message)
}

func (logger *WriterSystemLogger) Warning(message string) error {
	return logger.writeMessage("WARNING", message)
}

func (logger *WriterSystemLogger) Notice(message string) error {
	return logger.writeMessage("NOTICE", message)
}

func (logger *WriterSystemLogger) Info(message string) error {
	return logger.writeMessage("INFO", message)
}

func (logger *WriterSystemLogger) Debug(message string) error {
	return logger.writeMessage("DEBUG", message)
}

func CreateLogger(out string) (SystemLogger, error) {

	switch out {

	case "stdio":
		return NewWriterSystemLogger(os.Stdout), nil

	case "syslog":
		return syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, GetProgramName())
	}

	return nil, errors.New("Unknown logging implementation: " + out)
}
