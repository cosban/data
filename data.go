package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Statement struct {
	Query string
	Args  []interface{}
}

func Connect(connection string) {
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

func SetString(i interface{}, otherwise string) string {
	check := i.(sql.NullString)
	if check.Valid {
		return check.String
	}
	return otherwise
}

func SetInt(i interface{}, otherwise int) int {
	n := i.(sql.NullInt64)
	if n.Valid {
		return int(n.Int64)
	}
	return otherwise
}

func GetDateString(i interface{}) string {
	date := i.(time.Time)
	return fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())
}

func GetDateTimeString(i interface{}) string {
	date := i.(time.Time)
	hour, min, sec := date.Clock()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", date.Year(), date.Month(), date.Day(), hour, min, sec)
}

func Prepare(query string, args ...interface{}) Statement {
	return Statement{query, args}
}

func ExecuteTransaction(stmts ...Statement) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, v := range stmts {
		_, err = tx.Exec(v.Query, v.Args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func QueryRow(stmt Statement, data ...interface{}) error {
	row, err := Query(stmt)
	defer row.Close()
	if err != nil {
		log.Print(err)
		return err
	}
	if row.Next() {
		err = row.Scan(data...)
	}
	return err
}

func Query(stmt Statement) (*sql.Rows, error) {
	return db.Query(stmt.Query, stmt.Args...)
}
