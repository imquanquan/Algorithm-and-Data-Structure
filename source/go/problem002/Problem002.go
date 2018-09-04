package problem002

import "strings"

func ReplaceBlank(s string) string {
	return strings.Replace(s, " ", "20%", -1)
}