package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
	"errors"
	"strings"
)

func main() {

	mainConfig, err := NewMainConfig()

	if err != nil {
		log.Fatal(err)
	}

	err = mainConfig.LoadConfigFiles(GetConfigFilename(), GetConfigSearchPaths())

	if err != nil {
		log.Fatal(err)
	}

	err = mainConfig.LoadCommandLineArgs()

	if err != nil {
		log.Fatal(err)
	}

	logger, err := CreateLogger(mainConfig.Logging)

	if err != nil {
		log.Fatal(err)
	}

	db, err := OpenDatabaseConnection(mainConfig.Connection.String())
	
	if err != nil {
		logger.Alert(err.Error())
		log.Fatal(err)
	}

	defer db.Close()

	status, err := QuerySlaveStatus(db)

	if err != nil {
		logger.Alert(err.Error())
		log.Fatal(err)
	}

	slaveIORunning := status["Slave_IO_Running"]

	if slaveIORunning.Valid != true {
		err = errors.New("No value for Slave_IO_Running")
		logger.Alert(err.Error())
		log.Fatal(err)
	}

	var logStatus func(string)

	if strings.ToLower(slaveIORunning.String) == "Yes" {
		
		logStatus = func(message string) {
			logger.Info(message)
		}

	} else {

		logStatus = func(message string) {
			logger.Alert(message)		
		}
	}

	logStatus(strings.Join([]string{
		status["Slave_IO_Running"].String,
		status["Slave_SQL_Running"].String,
		status["Seconds_Behind_Master"].String,
		status["Read_Master_Log_Pos"].String,
		status["Relay_Log_Pos"].String,
		status["Exec_Master_Log_Pos"].String,
		status["Master_Log_File"].String,
		"\"" + status["Slave_IO_State"].String + "\"",
	}, " "))
}

func GetProgramName() string {
	return os.Args[0]
}

func GetConfigFilename() string {
	return GetProgramName() + ".json"
}

func GetConfigSearchPaths() []string {
	return []string{"~", "."}
}

func OpenDatabaseConnection(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, err
}
