package main

import (
	"encoding/json"
	"flag"
	"os"
	"os/user"
)

type MainConfig struct {
	Connection TCPDSN
	Logging    string
}

func NewMainConfig() (*MainConfig, error) {

	usr, err := user.Current()

	if err != nil {
		usr, err = user.Lookup("root")
		if err != nil {
			return nil, err
		}
	}

	return &MainConfig{
		Connection: TCPDSN{
			Hostname:  "127.0.0.1",
			Port:      3306,
			Username:  usr.Username,
			Password:  "",
			Collation: "utf8_general_ci",
		},
		Logging: "stdio",
	}, nil
}

func (config *MainConfig) loadConfigFile(filename string) error {

	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(config)

	if err != nil {
		return err
	}

	return nil
}

func (config *MainConfig) LoadConfigFiles(filename string, searchPaths []string) error {

	for _, filepath := range FindFile(filename, searchPaths) {

		err := config.loadConfigFile(filepath)

		if err != nil {
			return err
		}
	}

	return nil
}

func (config *MainConfig) LoadCommandLineArgs() error {

	flag.StringVar(&config.Connection.Hostname, "hostname", config.Connection.Hostname, "server hostname")
	flag.UintVar(&config.Connection.Port, "port", config.Connection.Port, "server port")

	flag.StringVar(&config.Connection.Username, "username", config.Connection.Username, "username")
	flag.StringVar(&config.Connection.Password, "password", config.Connection.Password, "password")

	flag.StringVar(&config.Connection.Database, "database", config.Connection.Database, "database")
	flag.StringVar(&config.Connection.Collation, "collation", config.Connection.Collation, "collation used")

	flag.StringVar(&config.Logging, "logging", config.Logging, "logging destination")

	flag.Parse()

	return nil
}
