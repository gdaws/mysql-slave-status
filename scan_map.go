package main

import "database/sql"

func ScanMap(rows * sql.Rows) (map[string]sql.NullString, error) {

	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		err = rows.Err()
		if err != nil {
			return nil, err
		} else {
			return nil, nil
		}
	}

	values := make([]interface{}, len(columns))

	for index := range values {
		values[index] = new(sql.NullString)
	}

	err = rows.Scan(values...)

	if err != nil {
		return nil, err
	}

	result := make(map[string]sql.NullString)

	for index, columnName := range columns {
		result[columnName] = *values[index].(*sql.NullString)
	}

	return result, nil
}
