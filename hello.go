package hello_go

import "strings"

func hello(input string) string {
	output := []string{"Hello", input + "!"}

	return strings.Join(output, " ")
}
