package multiplexer

func GetUser(user string) string {
	switch len(user) {
	case 22:
		return user[3:][:18]
	case 21:
		return user[2:][:18]
	default:
		return user
	}
}
