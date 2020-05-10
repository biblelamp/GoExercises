package util

import "strings"

func SplitString(str string) []string {
	var result []string
	part := ""
	isString := false
	for _, item := range str {
		c := string(item)
		if c == "\"" {
			isString = !isString
			if isString {
				if len(part) > 0 {
					result = append(result, part)
					part = ""
				}
				part += c
			} else {
				result = append(result, part+"\"")
				part = ""
			}
		} else if strings.Contains("!#:", c) {
			if len(part) > 0 {
				result = append(result, strings.TrimSpace(part))
			}
			result = append(result, c)
			part = ""
		} else if (c == ",") && !isString {
			if len(part) > 0 {
				result = append(result, strings.TrimSpace(part))
				part = ""
			}
		} else {
			part += c
		}
	}
	if len(part) > 0 {
		result = append(result, strings.TrimSpace(part))
	}
	return result
}
