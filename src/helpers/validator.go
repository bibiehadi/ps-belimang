package helpers

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

func ValidateUrl(url string) bool {
	match, err := regexp.MatchString(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`, url)
	if err != nil {
		fmt.Printf("error validating phone number: %", err)
		return false
	}
	if !match {
		fmt.Printf("invalid phone number format (must start with + and valid international code)")
		return false
	}
	return true
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
