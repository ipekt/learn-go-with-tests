package stringlib

import "strings"

func HasP(word string) bool {
	return strings.ContainsAny(word, "p")
}
