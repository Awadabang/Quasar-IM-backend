package util

import "regexp"

func IsSupportedPassword(password string) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}
