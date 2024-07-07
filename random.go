package random

import (
	"math/rand"
)

var (
	Numeric      = []rune("0123456789")
	AlphaLower   = []rune("abcdefghijklmnopqrstuvwxyz")
	AlphaUpper   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	RussianLower = []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
	RussianUpper = []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
	Special      = []rune("!@#$%^&*()_-+=[]{}\\|/?>.<,'\"`~")

	AlphaNumeric            = append(append(AlphaLower, AlphaUpper...), Numeric...)
	AlphaNumericWithSpecial = append(AlphaNumeric, Special...)
)

func String(set []rune, l uint) string {
	s := make([]rune, l)

	for i := range s {
		s[i] = set[rand.Intn(len(set))]
	}

	return string(s)
}
