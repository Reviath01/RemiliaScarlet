package multiplexer

func GetUser(user string) string {
	if len(user) == 18 {
		return user
	} else if len(user) == 22 {
		return user[3:][:18]
	} else if len(user) == 21 {
		return user[2:][:18]
	}
	return user
}
