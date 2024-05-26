package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:689599Df;@tcp(127.0.0.1:3306)/goProduct?charset=utf8")
	return
}

// Return one row of data
func GetResultRow(rows *sql.Rows) map[string]string {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	// save rows into record directory
	record := make(map[string]string)
	for rows.Next() {
		// fill values with the corresponding column data from the query result.
		rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				//fmt.Println(reflect.TypeOf(col))
				record[columns[i]] = string(v)
			}
		}
	}
	return record
}

// Return all results
func GetResultRows(rows *sql.Rows) map[int]map[string]string {
	// return all columns
	columns, _ := rows.Columns()

	// all values in one row for each column, which could be expressed in []byte
	vals := make([][]byte, len(columns))

	// a row of filled data
	scans := make([]interface{}, len(columns))

	// Fill data in vals
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		// fill in data
		rows.Scan(scans...)
		// every row
		row := make(map[string]string)
		// copy data in vals into row
		for k, v := range vals {
			key := columns[k]
			row[key] = string(v)
		}
		result[i] = row
		i++
	}
	return result
}
