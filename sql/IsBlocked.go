package sql

import "fmt"

// IsBlocked checks if the command is blocked, returns a string
func IsBlocked(guildid string, commandname string) string {
	db := Connect()
	defer db.Close()
	type Tag struct {
		isblocked string
	}

	var tag Tag

	err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", commandname, guildid)).Scan(&tag.isblocked)

	if err != nil {
		return "nil"
	}
	return "true"
}
