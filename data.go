package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Statement is a struct which represents a prepared statement.
// Calling Prepare allows the user to not have to construct this manually.
type Statement struct {
	Query string
	Args  []interface{}
}

// Connect takes in a valid psql connection string and opens the connection
func Connect(connection string) {
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

// SetString returns the value of i if it isn't null, otherwise returns otherwise
func SetString(i interface{}, otherwise string) string {
	check := i.(sql.NullString)
	if check.Valid {
		return check.String
	}
	return otherwise
}

// SetInt returns the value of i if it isn't null, otherwise returns otherwise
func SetInt(i interface{}, otherwise int) int {
	n := i.(sql.NullInt64)
	if n.Valid {
		return int(n.Int64)
	}
	return otherwise
}

// GetDateString returns a formatted date string
func GetDateString(i interface{}) string {
	date := i.(time.Time)
	return fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())
}

// GetDateTimeString returns a formatted date-time string
func GetDateTimeString(i interface{}) string {
	date := i.(time.Time)
	hour, min, sec := date.Clock()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", date.Year(), date.Month(), date.Day(), hour, min, sec)
}

// Prepare creates a Statement which may be used in any of the Execute or Query functions
func Prepare(query string, args ...interface{}) Statement {
	return Statement{query, args}
}

// PrepareAndExecute performs the actions of Prepare and Execute a single statement
// in one step rather than forcing the user to manually Prepare a statement.
func PrepareAndExecute(query string, args ...interface{}) error {
	return ExecuteTransaction(Statement{query, args})
}

// PrepareAndQuery performs the actions of Prepare and Query a single statement
// in one step rather than forcing the user to manually Prepare a statement.
func PrepareAndQuery(query string, args ...interface{}) (*sql.Rows, error) {
	return Query(Statement{query, args})
}

// ExecuteTransaction allows the user to execute >= 1 statement in a single sql transaction
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

// Query allows the user to perform a query where >= 1 rows are expected back
// It simply returns the rows so that the user may manipulate the data as they please
func Query(stmt Statement) (*sql.Rows, error) {
	return db.Query(stmt.Query, stmt.Args...)
}

// QueryRow allows the user to perform an sql query where only one row is expected
// The results of the query are put into the passed in data interfaces
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
