package sql

import "fmt"

//CheckLanguage function for all commands.
func CheckLanguage(guildid string) string {
	db := Connect()
	defer db.Close()
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
