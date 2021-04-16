package sql

import "fmt"

func IsBlocked(guildid string, commandname string) string {
	db := Connect()
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
