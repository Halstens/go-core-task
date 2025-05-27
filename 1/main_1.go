package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

type numberSystem struct {
	namePath    *int
	valueSystem string
}

func main() {

	var varsArray []interface{}

	salt := "go-2024"

	var numDecimal int = 42   // Десятичная система
	var numOctal int = 052    // Восьмеричная система
	var numHexadecimal = 0x2A // Шестнадцатиричная система

	listDecimal := numberSystem{
		&numDecimal,
		"dec",
	}
	listOctal := numberSystem{
		&numOctal,
		"octal",
	}
	listHex := numberSystem{
		&numHexadecimal,
		"hexadec",
	}

	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	varsArray = append(varsArray, listDecimal, listOctal, listHex,
		pi, name, isActive, complexNum)

	checkTypeVariables(varsArray)
	str := transTypeToString(varsArray)
	fmt.Println("Строка: ", str)

	strToRune := []rune(str)

	hash := makeSHA256(strToRune, salt)

	fmt.Println("Хэш: ", hash)

}

func checkTypeVariables(array []interface{}) {
	for _, item := range array {
		if ns, ok := item.(numberSystem); ok {
			fmt.Println(reflect.TypeOf(*ns.namePath))
		} else {
			fmt.Println(reflect.TypeOf(item))
		}
	}
}

func transTypeToString(array []interface{}) string {
	var builder strings.Builder
	for i, item := range array {
		if i > 0 {
			builder.WriteString(" | ")
		}
		if ns, ok := item.(numberSystem); ok {
			switch ns.valueSystem {
			case "octal":
				builder.WriteString("0")
				builder.WriteString(fmt.Sprintf("%o", *ns.namePath))
			case "hexadec":
				builder.WriteString("0x")
				builder.WriteString(fmt.Sprintf("%x", *ns.namePath))
			default:
				builder.WriteString(fmt.Sprint(*ns.namePath))
			}
		} else {
			builder.WriteString(fmt.Sprint(item))
		}

	}
	result := builder.String()
	return result
}

func makeSHA256(rn []rune, salt string) string {
	midRunes := len(rn) / 2
	rnSalt := []rune(salt)
	runesWithSalt := make([]rune, 0, len(rn)+len(rnSalt))
	runesWithSalt = append(runesWithSalt, rn[:midRunes]...)
	runesWithSalt = append(runesWithSalt, rnSalt...)
	runesWithSalt = append(runesWithSalt, rn[midRunes:]...)

	data := []byte(string(runesWithSalt))
	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])

}
