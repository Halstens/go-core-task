package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func prepareData() []interface{} {

	var varsArray []interface{}
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
	return varsArray
}

func TestTransTypeToString(t *testing.T) {
	exceptStr := "42 | 052 | 0x2a | 3.14 | Golang | true | (1+2i)"
	result := transTypeToString(prepareData())

	if result != exceptStr {
		t.Error(result, exceptStr)
	}
}

func TestCheckTypeVariables(t *testing.T) {
	expected := "int\nint\nint\nfloat64\nstring\nbool\ncomplex64\n"
	oldStd := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	checkTypeVariables(prepareData())
	w.Close()
	os.Stdout = oldStd

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if output != expected {
		t.Error(output, expected)
	}
}

func TestMakeSHA256(t *testing.T) {
	expected := "1d7810097f99223580df88c2fe6d3b7c67752f74032be5d87b713a64d81e64ce"

	result := makeSHA256([]rune("42 | 052 | 0x2a | 3.14 | Golang | true | (1+2i)"), "go-2024")

	if expected != result {
		t.Error(result, expected)
	}
}
