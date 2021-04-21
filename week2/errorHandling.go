package week2

import "database/sql"

func QueryDbData(db *sql.DB, querySql string) (rows int, err error) {
	return rows, sql.ErrNoRows
}
