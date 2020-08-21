package bigint

// BigInt represents a big integer as an array of small integers.
type BigInt struct {
	Digits []int
	IsNeg  bool
}

var maxDigits int
var zeroArray []int
var bigZero, bigOne BigInt

// SetMaxDigits initializes the RSA library.
func SetMaxDigits(value int) {
	if maxDigits != value {
		maxDigits = value
		zeroArray = make([]int, value, value)
		bigZero = NewBigInt(false)
		bigOne = NewBigInt(false)
		bigOne.Digits[0] = 1
	}
}

// NewBigInt initializes a new BigInt.
func NewBigInt(flag bool) BigInt {
	bi := BigInt{}
	if flag {
		bi.Digits = []int{}
	} else {
		// JS manually populates and copies zeroArray,
		// but we can simply make a new slice using the length of zeroArray
		bi.Digits = make([]int, len(zeroArray), len(zeroArray))
	}
	bi.IsNeg = false
	return bi
}

var BiRadixBits = 16
var BiRadix = 1 << 16 // = 2^16 = 65536
var BiHalfRadix = BiRadix >> 1
var BiRadixSquared = BiRadix * BiRadix

var bitsPerDigit = BiRadixBits
var maxDigitVal = BiRadix - 1

// func biFromDecimal(s string) BigInt {}

func BiCopy(bi BigInt) BigInt {
	result := NewBigInt(true)
	result.Digits = append(result.Digits, bi.Digits...)
	result.IsNeg = bi.IsNeg
	return result
}

// func biFromNumber(i int) BigInt {}

// func BiToString(x BigInt, radix int) string {}

// func biToDecimal(x BigInt) string {}

func BiToHex(x BigInt) string {
	var result string
	for i := BiHighIndex(x); i > -1; i-- {
		result += digitToHex(x.Digits[i])
	}
	return result
}

func BiFromHex(s string) BigInt {
	result := NewBigInt(false)

	// TODO: understand why we don't need to account for "-"
	if s[0] == "-"[0] {
		result.IsNeg = true
	}

	sl := len(s)
	for i, j := sl, 0; i > 0; i, j = i-4, j+1 {
		c := max(i-4, 0)
		result.Digits[j] = hexToDigit(string(s[c : c+min(i, 4)]))
	}
	return result
}

// func biFromString(s string, radix int) BigInt {}

func BiToBytes(x BigInt) string {
	var result string
	for i := BiHighIndex(x); i > -1; i-- {
		result += digitToBytes(x.Digits[i])
	}
	return result
}

// func biDump(b BigInt) string {}

func BiAdd(x BigInt, y BigInt) BigInt {
	if x.IsNeg != y.IsNeg {
		y.IsNeg = !y.IsNeg
		result := BiSubtract(x, y)
		y.IsNeg = !y.IsNeg
		return result
	}

	result := NewBigInt(false)

	c, n := 0, 0
	for i := range x.Digits {
		n = x.Digits[i] + y.Digits[i] + c
		result.Digits[i] = n & 0xFFFF
		c = 0
		if n >= BiRadix {
			c++
		}
	}

	// Check for the unbelievably stupid degenerate case of r == -0.
	if result.Digits[0] == 0 && BiHighIndex(result) == 0 {
		result.IsNeg = false
	} else {
		result.IsNeg = x.IsNeg
	}
	return result
}

func BiSubtract(x BigInt, y BigInt) BigInt {
	if x.IsNeg != y.IsNeg {
		y.IsNeg = !y.IsNeg
		result := BiAdd(x, y)
		y.IsNeg = !y.IsNeg
		return result
	}

	result := NewBigInt(false)

	c, n := 0, 0
	for i := range x.Digits {
		n = x.Digits[i] - y.Digits[i] + c
		result.Digits[i] = n & 0xFFFF
		if result.Digits[i] < 0 {
			result.Digits[i] += BiRadix
		}
		c = 0
		if n < 0 {
			c--
		}
	}

	if c == -1 {
		c = 0
		for i := range x.Digits {
			n = 0 - result.Digits[i] + c
			result.Digits[i] = n & 0xFFFF
			if result.Digits[i] < 0 {
				result.Digits[i] += BiRadix
			}
			c = 0
			if n < 0 {
				c--
			}
		}
		result.IsNeg = !x.IsNeg
	} else {
		result.IsNeg = x.IsNeg
	}

	// Check for the unbelievably stupid degenerate case of r == -0.
	if result.Digits[0] == 0 && BiHighIndex(result) == 0 {
		result.IsNeg = false
	}
	return result
}

func BiHighIndex(x BigInt) int {
	result := len(x.Digits) - 1
	for {
		if result == 0 || x.Digits[result] != 0 {
			break
		}
		result--
	}
	return result
}

func BiNumBits(x BigInt) int {
	n := BiHighIndex(x)
	d := x.Digits[n]
	m := (n + 1) * bitsPerDigit
	result := 0
	for result = m; result > m-bitsPerDigit; result-- {
		if d&0x8000 != 0 {
			break
		}
		d = d << 1
	}
	return result
}

func BiMultiply(x BigInt, y BigInt) BigInt {
	n := BiHighIndex(x)
	t := BiHighIndex(y)
	result := NewBigInt(false)

	for i := 0; i <= t; i++ {
		c, k, uv := 0, i, 0
		for j := 0; j <= n; j, k = j+1, k+1 {
			uv = result.Digits[k] + x.Digits[j]*y.Digits[i] + c
			result.Digits[k] = uv & maxDigitVal
			c = uv >> BiRadixBits
		}
		result.Digits[i+n+1] = c
	}
	result.IsNeg = x.IsNeg != y.IsNeg
	return result
}

func BiMultiplyDigit(x BigInt, y int) BigInt {
	var c, uv int
	n := BiHighIndex(x)
	result := NewBigInt(false)
	for j := 0; j <= n; j++ {
		uv = result.Digits[j] + x.Digits[j]*y + c
		result.Digits[j] = uv & maxDigitVal
		c = uv >> BiRadixBits
	}
	result.Digits[1+n] = c
	return result
}

func BiShiftLeft(x BigInt, n int) BigInt {
	digitCount := n / bitsPerDigit
	result := NewBigInt(false)
	arrayCopy(x.Digits, 0, result.Digits, digitCount, len(result.Digits)-digitCount)
	bits := n % bitsPerDigit
	rightBits := bitsPerDigit - bits
	for i := len(result.Digits) - 1; i > 0; i-- {
		result.Digits[i] = ((result.Digits[i] << bits) & maxDigitVal) | ((result.Digits[i-1] & highBitMasks[bits]) >> (rightBits))
	}
	result.Digits[0] = ((result.Digits[0] << bits) & maxDigitVal)
	result.IsNeg = x.IsNeg
	return result
}

func BiShiftRight(x BigInt, n int) BigInt {
	digitCount := n / bitsPerDigit
	result := NewBigInt(false)
	arrayCopy(x.Digits, digitCount, result.Digits, 0, len(x.Digits)-digitCount)
	bits := n % bitsPerDigit
	leftBits := bitsPerDigit - bits
	for i := 0; i < len(result.Digits)-1; i++ {
		result.Digits[i] = (result.Digits[i] >> bits) | ((result.Digits[i+1] & lowBitMasks[bits]) << leftBits)
	}
	result.Digits[len(result.Digits)-1] >>= bits
	return result
}

func BiMultiplyByRadixPower(x BigInt, n int) BigInt {
	result := NewBigInt(false)
	arrayCopy(x.Digits, 0, result.Digits, n, len(result.Digits)-n)
	return result
}

func BiDivideByRadixPower(x BigInt, n int) BigInt {
	result := NewBigInt(false)
	arrayCopy(x.Digits, n, result.Digits, 0, len(result.Digits)-n)
	return result
}

func BiModuloByRadixPower(x BigInt, n int) BigInt {
	result := NewBigInt(false)
	arrayCopy(x.Digits, 0, result.Digits, 0, n)
	return result
}

func BiCompare(x BigInt, y BigInt) int {
	if x.IsNeg != y.IsNeg {
		if x.IsNeg {
			// neg x is less than pos y
			return -1
		}
		// pos x is greater than neg y
		return 1
	}

	l := len(x.Digits) - 1
	for i := range x.Digits {
		j := l - i
		if x.Digits[j] != y.Digits[j] {
			if x.IsNeg {
				if x.Digits[j] > y.Digits[j] {
					// neg x is less than neg y
					return -1
				}
				// neg x is greater than neg y
				return 1
			}
			if x.Digits[j] < y.Digits[j] {
				// pos x is less than pos y
				return -1
			}
			// pos x is greater than pos y
			return 1
		}
	}
	// x is equal to y
	return 0
}

func BiDivideModulo(x BigInt, y BigInt) [2]BigInt {
	nb := BiNumBits(x)
	tb := BiNumBits(y)
	origYIsNeg := y.IsNeg
	var q, r BigInt

	if nb < tb {
		// |x| < |y|
		if x.IsNeg {
			q = BiCopy(bigOne)
			q.IsNeg = !y.IsNeg
			x.IsNeg = false
			y.IsNeg = false
			r = BiSubtract(y, x)
			// Restore signs, 'cause they're references.
			x.IsNeg = true
			y.IsNeg = origYIsNeg
		} else {
			q = NewBigInt(false)
			r = BiCopy(x)
		}
		return [2]BigInt{q, r}
	}

	q = NewBigInt(false)
	r = x

	// Normalize Y.
	t := ceil(tb, bitsPerDigit) - 1
	lambda := 0
	for y.Digits[t] < BiHalfRadix {
		y = BiShiftLeft(y, 1)
		lambda++
		tb++
		t = ceil(tb, bitsPerDigit) - 1
	}

	// Shift r over to keep the quotient constant. We'll shift the
	// remainder back at the end.
	r = BiShiftLeft(r, lambda)
	nb += lambda // Update the bit count for x.
	n := ceil(nb, bitsPerDigit) - 1

	b := BiMultiplyByRadixPower(y, n-t)
	for BiCompare(r, b) != -1 {
		q.Digits[n-t]++
		r = BiSubtract(r, b)
	}

	i := n
	for i = n; i > t; i-- {
		var ri, ri1, ri2, yt, yt1 int
		if !(i >= len(r.Digits)) {
			ri = r.Digits[i]
		}
		if !(i-1 >= len(r.Digits)) {
			ri1 = r.Digits[i-1]
		}
		if !(i-2 >= len(r.Digits)) {
			ri2 = r.Digits[i-2]
		}
		if !(t >= len(y.Digits)) {
			yt = y.Digits[t]
		}
		if !(t-1 >= len(y.Digits)) {
			yt1 = y.Digits[t-1]
		}
		if ri == yt {
			q.Digits[i-t-1] = maxDigitVal
		} else {
			q.Digits[i-t-1] = (ri*BiRadix + ri1) / yt
		}

		c1 := q.Digits[i-t-1] * ((yt * BiRadix) + yt1)
		c2 := (ri * BiRadixSquared) + ((ri1 * BiRadix) + ri2)
		for c1 > c2 {
			q.Digits[i-t-1]--
			c1 = q.Digits[i-t-1] * ((yt * BiRadix) | yt1)
			c2 = (ri * BiRadix * BiRadix) + ((ri1 * BiRadix) + ri2)
		}

		b = BiMultiplyByRadixPower(y, i-t-1)
		r = BiSubtract(r, BiMultiplyDigit(b, q.Digits[i-t-1]))
		if r.IsNeg {
			r = BiAdd(r, b)
			q.Digits[i-t-1]--
		}
	}

	r = BiShiftRight(r, lambda)
	// Fiddle with the signs and stuff to make sure that 0 <= r < y.
	q.IsNeg = x.IsNeg != origYIsNeg
	if x.IsNeg {
		if origYIsNeg {
			q = BiAdd(q, bigOne)
		} else {
			q = BiSubtract(q, bigOne)
		}
		y = BiShiftRight(y, lambda)
		r = BiSubtract(y, r)
	}
	// Check for the unbelievably stupid degenerate case of r == -0.
	if r.Digits[0] == 0 && BiHighIndex(r) == 0 {
		r.IsNeg = false
	}

	return [2]BigInt{q, r}
}

func BiDivide(x BigInt, y BigInt) BigInt {
	return BiDivideModulo(x, y)[0]
}

func BiModulo(x BigInt, y BigInt) BigInt {
	return BiDivideModulo(x, y)[1]
}

// func biMultiplyMod(x bigInt, y bigInt, m bigInt) bigInt {}

// func biPow(x BigInt, y BigInt) BigInt {}

// func biPowMod(x BigInt, y BigInt, m BigInt) BigInt {}
