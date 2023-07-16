package utils

import "regexp"

func IsEmailValid(email string) bool {
	regex_pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regexExp := regexp.MustCompile(regex_pattern)
	return regexExp.MatchString(email)
}
