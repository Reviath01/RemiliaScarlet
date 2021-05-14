package multiplexer

// GetRole function provides to get a role with id or mention
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
