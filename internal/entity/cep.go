package entity

import (
	"regexp"
	"strings"
)

type (
	CEP string
)

func (c CEP) IsValid() bool {
	matches, _ := regexp.Match(`^\d{5}-?\d{3}$`, []byte(c))
	return matches
}

func (c CEP) GetDigits() string {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(string(c), -1)
	return strings.Join(matches, "")
}
