package util

import (
	"regexp"
	"strings"
)

func IsValidGroupNumber(numGrp string) bool {
	result, _ := regexp.MatchString(`^[1-9]{1}\d?`, numGrp)
	return result
}

func IsValidLineNumber(numLine string) bool {
	result, _ := regexp.MatchString(`^[1-9]{1}\d?\.\d{1}\d?$`, numLine)
	return result
}

func SplitStringWithDelimiter(str, delimiter string) []string {
	var result []string
	if !strings.Contains(str, delimiter) {
		result = append(result, str)
		return result
	}
	part := ""
	isString := false
	for _, item := range str {
		c := string(item)
		if c == "\"" {
			isString = !isString
		} else if (c == delimiter) && !isString {
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
