package main

import (
	"database/sql"
)

func QuerySlaveStatus(db *sql.DB) (map[string]sql.NullString, error) {

	rows, err := db.Query("SHOW SLAVE STATUS")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	status, err := ScanMap(rows)

	if err != nil {
		return nil, err
	}

	return status, nil
}
