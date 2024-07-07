package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/teran/go-random"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("USAGE: %s SET LENGTH\n\n", os.Args[0])
		fmt.Println("SET - is name of supported character set, multiple sets can be specified as a + separated list:")
		fmt.Println("* numeric")
		fmt.Println("* alpha_lower")
		fmt.Println("* alpha_upper")
		fmt.Println("* russian_lower")
		fmt.Println("* russian_upper")
		fmt.Println("* special")
		fmt.Println("* alphanumeric_with_special")
		fmt.Print("* alphanumeric\n\n")

		fmt.Print("LENGTH - the generated word length\n\n")

		os.Exit(1)
	}

	charList, err := characterSetByName(os.Args[1])
	if err != nil {
		panic(err)
	}

	length, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(random.String(charList, uint(length)))
}

func characterSetByName(name string) ([]rune, error) {
	sets := strings.Split(name, "+")
	charList := []rune{}

	for _, set := range sets {
		switch set {
		case "numeric":
			charList = append(charList, random.Numeric...)
		case "alpha_lower":
			charList = append(charList, random.AlphaLower...)
		case "alpha_upper":
			charList = append(charList, random.AlphaUpper...)
		case "russian_lower":
			charList = append(charList, random.RussianLower...)
		case "russian_upper":
			charList = append(charList, random.RussianUpper...)
		case "special":
			charList = append(charList, random.Special...)
		case "alphanumeric_with_special":
			charList = append(charList, random.AlphaNumericWithSpecial...)
		case "alphanumeric":
			charList = append(charList, random.AlphaNumeric...)
		default:
			return nil, errors.Errorf("unexpected set name: `%s`", set)
		}
	}

	return charList, nil
}
