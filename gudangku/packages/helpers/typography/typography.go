package typography

import "strings"

func UcFirst(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

func RemoveLastChar(str, char string) string {
	word := strings.ToLower(str)
	character := strings.ToLower(char)

	if len(word) > 0 && strings.HasSuffix(word, character) {
		return word[:len(word)-1]
	}
	return word
}
