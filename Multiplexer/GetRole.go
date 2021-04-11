package multiplexer

func GetRole(role string) string {
	if len(role) == 18 {
		return role
	} else if len(role) == 22 {
		return role[3:][:18]
	}
	return "Error"
}
