package validator

import (
	"regexp"
)

func CheckLength(data string, maxSize int) bool {
	if len(data) > maxSize {
		return false
	}
	return true
}
func CheckPhone(data string) bool {
	regex := "(0[3|5|7|8|9])+([0-9]{8})\\b"
	match, _ := regexp.MatchString(regex, data)
	return match
}
func CheckMail(data string) bool {
	if !CheckLength(data,50) {
		return false
	}
	regex := "^[a-z][a-z0-9_\\.]{4,32}@[a-z0-9]{2,}(\\.[a-z0-9]{2,4}){1,2}$"
	match, _ := regexp.MatchString(regex, data)
	return match
}
