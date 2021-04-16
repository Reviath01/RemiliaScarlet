package sql

import "fmt"

func CheckLanguage(guildid string) string {
	db := Connect()
	type Tag struct {
		lang string
	}

	var tag Tag

	err := db.QueryRow(fmt.Sprintf("SELECT language FROM languages WHERE guildid ='%s'", guildid)).Scan(&tag.lang)

	if err != nil {
		return "nil"
	}
	return tag.lang
}
