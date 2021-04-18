package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	ReadConfig()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", User, Password, Host, Database))
	if err != nil {
		fmt.Printf("Cannot connect MySQL database %s\n", err.Error())
		return nil
	}
	return db
}
