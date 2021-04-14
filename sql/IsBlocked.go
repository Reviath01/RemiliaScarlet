package sql

func IsBlocked(guildid string, commandname string) string {
	db := Connect()
	type Tag struct {
		isblocked string
	}

	var tag Tag

	err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + commandname + "' AND guildid ='" + guildid + "'").Scan(&tag.isblocked)

	if err != nil {
		return "nil"
	} else {
		return "true"
	}
}
