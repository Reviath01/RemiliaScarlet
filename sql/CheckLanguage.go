package sql

func CheckLanguage(guildid string) string {
	db := Connect()
	type Tag struct {
		lang string
	}

	var tag Tag

	err := db.QueryRow("SELECT language FROM languages WHERE guildid ='" + guildid + "'").Scan(&tag.lang)

	if err != nil {
		return "nil"
	}
	return tag.lang
}
