package entity

import "regexp"

type (
	CEP string
)

func (c CEP) IsValid() bool {
	matches, _ := regexp.Match(`^\d{5}-?\d{3}$`, []byte(c))
	return matches
}

func (c CEP) GetDigits() string {
	re := regexp.MustCompile(`\d+`)
	return re.FindString(string(c))
}
