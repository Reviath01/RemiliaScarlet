package multiplexer

//Get role function for all commands.
func GetRole(role string) string {
	switch len(role) {
	case 18:
		return role
	case 22:
		return role[3:][:18]
	default:
		return "Error"
	}
}
