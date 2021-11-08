package utils

import "strings"

func ContainesName(str1 string, str2 string) bool {
	str1_WithoutSpace := ignoreSpaces(str1)
	str2_WithoutSpace := ignoreSpaces(str2)
	return strings.Contains(strings.ToLower(str1_WithoutSpace), strings.ToLower(str2_WithoutSpace))
}

func ignoreSpaces(str string) string {
	components := strings.Split(str, " ")
	var result string
	for _, component := range components {
		result += component
	}
	return result
}
