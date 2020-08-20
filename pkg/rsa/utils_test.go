package rsa

import (
	"testing"
)

func TestNewBigInt(t *testing.T) {
	SetMaxDigits(8)
	expected := []int{0, 0, 0, 0, 0, 0, 0, 0}
	result := NewBigInt(false)

	if len(expected) != len(result.Digits) {
		t.Errorf("NewBigInt(false) failed, expected %v, got %v", len(expected), len(result.Digits))
	}

	for i, e := range expected {
		if e != result.Digits[i] {
			t.Errorf("NewBigInt(false) failed, expected %v, got %v", expected, result.Digits)
		}
	}
}

func TestReverseStr(t *testing.T) {
	expected := "!dlroW olleH"
	result := reverseStr("Hello World!")

	if expected != result {
		t.Errorf("reverseStr(\"Hello World!\") failed, expected %v, got %v", expected, result)
	}
}

// digitToHex

func TestCharToHex(t *testing.T) {
	expected := []int{13, 14, 10, 13, 11, 14, 14, 15}
	result := []int{}
	for _, s := range "deadbeef" {
		result = append(result, charToHex(rune(s)))
	}

	for i, e := range expected {
		if e != result[i] {
			t.Errorf("charToHex(\"deadbeef\") failed, expected %v, got %v", expected, result)
			break
		}
	}
}

func TestHexToDigit(t *testing.T) {
	expected := []int{57005, 48879}
	result := []int{}
	for _, s := range []string{"dead", "beef"} {
		result = append(result, hexToDigit(s))
	}

	for i, e := range expected {
		if e != result[i] {
			t.Errorf("hexToDigit(\"deadbeef\") failed, expected %v, got %v", expected, result)
			break
		}
	}
}

// digitToBytes

func TestArrayCopy(t *testing.T) {
	src := []int{13, 14, 10, 13, 11, 14, 14, 15}  // deadbeef
	expected := []int{0, 0, 11, 14, 10, 13, 0, 0} // 00bead00
	result := []int{0, 0, 0, 0, 0, 0, 0, 0}

	arrayCopy(src, 2, result, 4, 2) // 0000ad00
	arrayCopy(src, 4, result, 2, 2) // 00be0000

	for i, e := range expected {
		if e != result[i] {
			t.Errorf("hexToDigit(\"deadbeef\") failed, expected %v, got %v", expected, result)
			break
		}
	}
}

// max

// min
