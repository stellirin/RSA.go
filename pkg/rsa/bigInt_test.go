package rsa

import (
	"testing"
)

// func TestBiFromDecimal(t *testing.T) {}

func TestBiCopy(t *testing.T) {
	SetMaxDigits(8)
	expected := biFromHex("deadbeef")
	result := biCopy(expected)

	for i, e := range expected.digits {
		if e != result.digits[i] {
			t.Errorf("biFromHex(\"deadbeef\") failed, expected %v, got %v", expected.digits, result.digits)
			break
		}
	}
	if expected.isNeg != result.isNeg {
		t.Errorf("biFromHex(\"deadbeef\") failed, expected isNeg %v, got %v", expected.isNeg, result.isNeg)
	}
}

// func TestBiFromNumber(t *testing.T) {}

// func TestBiToString(t *testing.T) {}

// func TestBiToDecimal(t *testing.T) {}

// func TestBiToHex(t *testing.T) {}

func TestBiFromHex(t *testing.T) {
	var expected, result bigInt
	SetMaxDigits(8)

	expected = bigInt{
		digits: []int{48879, 57005, 0, 0, 0, 0, 0, 0}, // beef, dead
		isNeg:  false,
	}
	result = biFromHex("deadbeef")
	for i, e := range expected.digits {
		if e != result.digits[i] {
			t.Errorf("biFromHex(\"deadbeef\") failed, expected %v, got %v", expected.digits, result.digits)
			break
		}
	}
	if expected.isNeg != result.isNeg {
		t.Errorf("biFromHex(\"deadbeef\") failed, expected isNeg %v, got %v", expected.isNeg, result.isNeg)
	}
}

func TestBiFromHexNeg(t *testing.T) {
	var expected, result bigInt
	SetMaxDigits(8)

	expected = bigInt{
		digits: []int{48879, 57005, 0, 0, 0, 0, 0, 0}, // beef, dead
		isNeg:  true,
	}
	result = biFromHex("-deadbeef")
	for i, e := range expected.digits {
		if e != result.digits[i] {
			t.Errorf("biFromHex(\"-deadbeef\") failed, expected %v, got %v", expected.digits, result.digits)
			// break
		}
	}
	if expected.isNeg != result.isNeg {
		t.Errorf("biFromHex(\"-deadbeef\") failed, expected isNeg %v, got %v", expected.isNeg, result.isNeg)
	}
}

// func TestBiFromString(t *testing.T) {}

// func TestBiToBytes(t *testing.T) {}

// func TestBiDump(t *testing.T) {}

func TestBiAdd(t *testing.T) {
	SetMaxDigits(8)
	bix := biFromHex("deadbeef")
	biy := biFromHex("beefdead")
	expected := []int{40348, 40349, 1, 0, 0, 0, 0, 0}
	result := biAdd(bix, biy)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biAdd(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			// break
		}
	}
}

func TestBiSubtract(t *testing.T) {
	SetMaxDigits(8)
	bix := biFromHex("deadbeef")
	biy := biFromHex("beefdead")
	expected := []int{57410, 8125, 0, 0, 0, 0, 0, 0}
	result := biSubtract(bix, biy)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biSubtract(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			// break
		}
	}
}

func TestBiHighIndex(t *testing.T) {
	SetMaxDigits(8)
	bi := biFromHex("deadbeef")
	expected := 1 // 0 index, every 4 chars
	result := biHighIndex(bi)

	if expected != result {
		t.Errorf("biHighIndex(BigInt) failed, expected %v, got %v", expected, result)
	}
}

func TestBiNumBits(t *testing.T) {
	SetMaxDigits(8)
	bi := biFromHex("deadbeef")
	expected := 32
	result := biNumBits(bi)

	if result != expected {
		t.Errorf("biNumBits(\"deadbeef\") failed, expected %v, got %v", expected, result)
	}
}

func TestBiMultiply(t *testing.T) {
	SetMaxDigits(8)
	bi := biFromHex("deadbeef")
	expected := []int{41761, 8557, 52498, 49585} // a321 216d cd12 c1b1
	result := biMultiply(bi, bi)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biMultiply(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			break
		}
	}
}

func TestBiMultiplyDigit(t *testing.T) {
	SetMaxDigits(8)
	bi := biFromHex("deadbeef")
	expected := []int{61168, 60123, 13, 0, 0, 0, 0, 0} //eef0, eadb, d
	result := biMultiplyDigit(bi, 16)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biMultiplyDigit(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			// break
		}
	}
}

func TestBiShiftLeft(t *testing.T) {
	SetMaxDigits(8)
	bi := biFromHex("deadbeef")
	expected := []int{0, 48879, 57005, 0, 0, 0, 0} // 0000, beef, dead
	result := biShiftLeft(bi, 16)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biShiftLeft(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			break
		}
	}
}

func TestBiShiftRight(t *testing.T) {
	SetMaxDigits(8)
	bi := biFromHex("deadbeef")
	expected := []int{57005, 0, 0, 0, 0, 0, 0} // dead
	result := biShiftRight(bi, 16)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biShiftLeft(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			break
		}
	}
}

// func TestBiMultiplyByRadixPower(t *testing.T) {}

func TestBiDivideByRadixPower(t *testing.T) {
	SetMaxDigits(8)
	bi := biFromHex("deadbeefcafef00dbeadfacebabe")
	expected := []int{61453, 51966, 48879, 57005, 0, 0, 0, 0} // f00d, cafe, beef, dead
	result := biDivideByRadixPower(bi, 3)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biDivideByRadixPower(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			break
		}
	}
}

func TestBiModuloByRadixPower(t *testing.T) {
	SetMaxDigits(8)
	bi := biFromHex("deadbeefcafef00dbeadfacebabe")
	expected := []int{47806, 64206, 48813, 61453, 0, 0, 0, 0} // babe, face, bead, f00d
	result := biModuloByRadixPower(bi, 4)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biModuloByRadixPower(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			break
		}
	}
}

func TestBiCompare(t *testing.T) {
	var result int

	w, x, y, z := "-bead", "-face", "dead", "beef"
	biw, bix, biy, biz := biFromHex(w), biFromHex(x), biFromHex(y), biFromHex(z)

	// neg x is less than pos y
	result = biCompare(bix, biy)
	if result != -1 {
		t.Errorf("biCompare(\"%v\", \"%v\") failed, expected %v, got %v", x, y, -1, result)
	}

	// pos x is greater than neg y
	result = biCompare(biz, bix)
	if result != 1 {
		t.Errorf("biCompare(\"%v\", \"%v\") failed, expected %v, got %v", z, x, 1, result)
	}

	// neg x is less than neg y
	result = biCompare(bix, biw)
	if result != -1 {
		t.Errorf("biCompare(\"%v\", \"%v\") failed, expected %v, got %v", x, w, -1, result)
	}

	// neg x is greater than neg y
	result = biCompare(biw, bix)
	if result != 1 {
		t.Errorf("biCompare(\"%v\", \"%v\") failed, expected %v, got %v", w, x, 1, result)
	}

	// pos x is less than pos y
	result = biCompare(biz, biy)
	if result != -1 {
		t.Errorf("biCompare(\"%v\", \"%v\") failed, expected %v, got %v", z, y, -1, result)
	}

	// pos x is greater than pos y
	result = biCompare(biy, biz)
	if result != 1 {
		t.Errorf("biCompare(\"%v\", \"%v\") failed, expected %v, got %v", y, z, 1, result)
	}

	// x is equal to y
	result = biCompare(biz, biz)
	if result != 0 {
		t.Errorf("biCompare(\"%v\", \"%v\") failed, expected %v, got %v", z, z, 0, result)
	}
}

func TestBiDivideModulo(t *testing.T) {
	SetMaxDigits(8)
	bi1 := biFromHex("deadbeef")
	bi2 := biFromHex("beefdead")
	expected := []int{1, 0, 0, 0, 0, 0, 0} // 1
	result := biDivideModulo(bi1, bi2)

	for i, e := range expected {
		if e != result[0].digits[i] {
			t.Errorf("biDivideModulo(\"deadbeef\") failed, expected %v, got %v, %v", expected, result[0].digits, result[1].digits)
			break
		}
	}
}

func TestBiDivide(t *testing.T) {
	SetMaxDigits(8)
	bi1 := biFromHex("deadbeef")
	bi2 := biFromHex("beefdead")
	expected := []int{1, 0, 0, 0, 0, 0, 0} // 1
	result := biDivide(bi1, bi2)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biDivideModulo(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			break
		}
	}
}

func TestBiModulo(t *testing.T) {
	SetMaxDigits(8)
	bi1 := biFromHex("deadbeef")
	bi2 := biFromHex("beefdead")
	expected := []int{57410, 8125, 0, 0, 0, 0, 0} // e042, 1fbd
	result := biModulo(bi1, bi2)

	for i, e := range expected {
		if e != result.digits[i] {
			t.Errorf("biDivideModulo(\"deadbeef\") failed, expected %v, got %v", expected, result.digits)
			break
		}
	}
}

// func TestBiMultiplyMod(t *testing.T) {}

// func TestBiPow(t *testing.T) {}

// func TestBiPowMod(t *testing.T) {}
