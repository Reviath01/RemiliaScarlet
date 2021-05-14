package multiplexer

import "strings"

//Getting message args function.
func GetArgs(message string, prefix string) []string {
	args := strings.Split(strings.TrimPrefix(message, prefix), " ")
	arguments, _ := Shift(args, 0)
	return arguments
}

//A required function for getting args.
func Shift(a []string, i int) ([]string, string) {
	b := a[i]
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a, b
}
