package main

import "fmt"

type TCPDSN struct {
	Hostname  string
	Port      uint
	Username  string
	Password  string
	Database  string
	Collation string
}

func (dsn *TCPDSN) String() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?collation=%s",
		dsn.Username,
		dsn.Password,
		dsn.Hostname,
		dsn.Port,
		dsn.Database,
		dsn.Collation)
}
