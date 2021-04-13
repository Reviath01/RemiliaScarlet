package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")
	if err != nil {
		fmt.Print("Cannot connect MySQL database " + err.Error())
		return nil
	}
	return db
}
