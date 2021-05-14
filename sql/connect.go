package sql

import (
	"database/sql"
	"fmt"

	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	_ "github.com/go-sql-driver/mysql"
)

// Connect returns to MySQL database
func Connect() *sql.DB {
	config.ReadConfig()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", config.User, config.Password, config.Host, config.Database))
	if err != nil {
		fmt.Printf("Cannot connect MySQL database %s\n", err.Error())
		return nil
	}
	return db
}
