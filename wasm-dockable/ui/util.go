package ui

import (
	"fmt"
	"regexp"
	"strings"
)

var upperRegex = regexp.MustCompile(`[[:upper:]]`)

func InlineStyle(styles ...map[string]string) string {
	var s []string
	for _, style := range styles {
		for k, v := range style {
			s = append(s, fmt.Sprintf("%s: %s;", upperRegex.ReplaceAllStringFunc(k, func(s string) string {
				return "-" + strings.ToLower(s)
			}), v))
		}
	}
	return strings.Join(s, " ")
}
