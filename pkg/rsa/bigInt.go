package rsa

var biRadixBits = 16
var bitsPerDigit = biRadixBits
var biRadix = 1 << 16 // = 2^16 = 65536
var biHalfRadix = biRadix >> 1
var biRadixSquared = biRadix * biRadix
var maxDigitVal = biRadix - 1

// func biFromDecimal(s string) BigInt {}

func biCopy(bi bigInt) bigInt {
	result := newBigInt(true)
	result.digits = append(result.digits, bi.digits...)
	result.isNeg = bi.isNeg
	return result
}

// func biFromNumber(i int) BigInt {}

func biToString(x bigInt, radix int) string {
	b := newBigInt(false)
	b.digits[0] = radix
	qr := biDivideModulo(x, b)
	result := hexatrigesimalToChar[qr[1].digits[0]]
	for biCompare(qr[0], bigZero) == 1 {
		qr = biDivideModulo(qr[0], b)
		result += hexatrigesimalToChar[qr[1].digits[0]]
	}
	var isNeg string
	if x.isNeg {
		isNeg = "-"
	}
	return isNeg + reverseStr(result)
}

// func biToDecimal(x BigInt) string {}

func biToHex(x bigInt) string {
	var result string
	for i := biHighIndex(x); i > -1; i-- {
		result += digitToHex(x.digits[i])
	}
	return result
}

func biFromHex(s string) bigInt {
	result := newBigInt(false)

	// TODO: understand why we don't need to account for "-"
	if s[0] == "-"[0] {
		result.isNeg = true
	}

	sl := len(s)
	for i, j := sl, 0; i > 0; i, j = i-4, j+1 {
		c := max(i-4, 0)
		result.digits[j] = hexToDigit(string(s[c : c+min(i, 4)]))
	}
	return result
}

// func biFromString(s string, radix int) BigInt {}

func biToBytes(x bigInt) string {
	var result string
	for i := biHighIndex(x); i > -1; i-- {
		result += digitToBytes(x.digits[i])
	}
	return result
}

// func biDump(b BigInt) string {}

func biAdd(x bigInt, y bigInt) bigInt {
	if x.isNeg != y.isNeg {
		y.isNeg = !y.isNeg
		result := biSubtract(x, y)
		y.isNeg = !y.isNeg
		return result
	}

	result := newBigInt(false)

	c, n := 0, 0
	for i := range x.digits {
		n = x.digits[i] + y.digits[i] + c
		result.digits[i] = n & 0xFFFF
		c = 0
		if n >= biRadix {
			c++
		}
	}

	// Check for the unbelievably stupid degenerate case of r == -0.
	if result.digits[0] == 0 && biHighIndex(result) == 0 {
		result.isNeg = false
	} else {
		result.isNeg = x.isNeg
	}
	return result
}

func biSubtract(x bigInt, y bigInt) bigInt {
	if x.isNeg != y.isNeg {
		y.isNeg = !y.isNeg
		result := biAdd(x, y)
		y.isNeg = !y.isNeg
		return result
	}

	result := newBigInt(false)

	c, n := 0, 0
	for i := range x.digits {
		n = x.digits[i] - y.digits[i] + c
		result.digits[i] = n & 0xFFFF
		if result.digits[i] < 0 {
			result.digits[i] += biRadix
		}
		c = 0
		if n < 0 {
			c--
		}
	}

	if c == -1 {
		c = 0
		for i := range x.digits {
			n = 0 - result.digits[i] + c
			result.digits[i] = n & 0xFFFF
			if result.digits[i] < 0 {
				result.digits[i] += biRadix
			}
			c = 0
			if n < 0 {
				c--
			}
		}
		result.isNeg = !x.isNeg
	} else {
		result.isNeg = x.isNeg
	}

	// Check for the unbelievably stupid degenerate case of r == -0.
	if result.digits[0] == 0 && biHighIndex(result) == 0 {
		result.isNeg = false
	}
	return result
}

func biHighIndex(x bigInt) int {
	result := len(x.digits) - 1
	for {
		if result == 0 || x.digits[result] != 0 {
			break
		}
		result--
	}
	return result
}

func biNumBits(x bigInt) int {
	n := biHighIndex(x)
	d := x.digits[n]
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

func biMultiply(x bigInt, y bigInt) bigInt {
	n := biHighIndex(x)
	t := biHighIndex(y)
	result := newBigInt(false)

	for i := 0; i <= t; i++ {
		c, k, uv := 0, i, 0
		for j := 0; j <= n; j, k = j+1, k+1 {
			uv = result.digits[k] + x.digits[j]*y.digits[i] + c
			result.digits[k] = uv & maxDigitVal
			c = uv >> biRadixBits
		}
		result.digits[i+n+1] = c
	}
	result.isNeg = x.isNeg != y.isNeg
	return result
}

func biMultiplyDigit(x bigInt, y int) bigInt {
	var c, uv int
	n := biHighIndex(x)
	result := newBigInt(false)
	for j := 0; j <= n; j++ {
		uv = result.digits[j] + x.digits[j]*y + c
		result.digits[j] = uv & maxDigitVal
		c = uv >> biRadixBits
	}
	result.digits[1+n] = c
	return result
}

func biShiftLeft(x bigInt, n int) bigInt {
	digitCount := n / bitsPerDigit
	result := newBigInt(false)
	arrayCopy(x.digits, 0, result.digits, digitCount, len(result.digits)-digitCount)
	bits := n % bitsPerDigit
	rightBits := bitsPerDigit - bits
	for i := len(result.digits) - 1; i > 0; i-- {
		result.digits[i] = ((result.digits[i] << bits) & maxDigitVal) | ((result.digits[i-1] & highBitMasks[bits]) >> (rightBits))
	}
	result.digits[0] = ((result.digits[0] << bits) & maxDigitVal)
	result.isNeg = x.isNeg
	return result
}

func biShiftRight(x bigInt, n int) bigInt {
	digitCount := n / bitsPerDigit
	result := newBigInt(false)
	arrayCopy(x.digits, digitCount, result.digits, 0, len(x.digits)-digitCount)
	bits := n % bitsPerDigit
	leftBits := bitsPerDigit - bits
	for i := 0; i < len(result.digits)-1; i++ {
		result.digits[i] = (result.digits[i] >> bits) | ((result.digits[i+1] & lowBitMasks[bits]) << leftBits)
	}
	result.digits[len(result.digits)-1] >>= bits
	return result
}

func biMultiplyByRadixPower(x bigInt, n int) bigInt {
	result := newBigInt(false)
	arrayCopy(x.digits, 0, result.digits, n, len(result.digits)-n)
	return result
}

func biDivideByRadixPower(x bigInt, n int) bigInt {
	result := newBigInt(false)
	arrayCopy(x.digits, n, result.digits, 0, len(result.digits)-n)
	return result
}

func biModuloByRadixPower(x bigInt, n int) bigInt {
	result := newBigInt(false)
	arrayCopy(x.digits, 0, result.digits, 0, n)
	return result
}

func biCompare(x bigInt, y bigInt) int {
	if x.isNeg != y.isNeg {
		if x.isNeg {
			// neg x is less than pos y
			return -1
		}
		// pos x is greater than neg y
		return 1
	}

	l := len(x.digits) - 1
	for i := range x.digits {
		j := l - i
		if x.digits[j] != y.digits[j] {
			if x.isNeg {
				if x.digits[j] > y.digits[j] {
					// neg x is less than neg y
					return -1
				}
				// neg x is greater than neg y
				return 1
			}
			if x.digits[j] < y.digits[j] {
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

func biDivideModulo(x bigInt, y bigInt) [2]bigInt {
	nb := biNumBits(x)
	tb := biNumBits(y)
	origYIsNeg := y.isNeg
	var q, r bigInt

	if nb < tb {
		// |x| < |y|
		if x.isNeg {
			q = biCopy(bigOne)
			q.isNeg = !y.isNeg
			x.isNeg = false
			y.isNeg = false
			r = biSubtract(y, x)
			// Restore signs, 'cause they're references.
			x.isNeg = true
			y.isNeg = origYIsNeg
		} else {
			q = newBigInt(false)
			r = biCopy(x)
		}
		return [2]bigInt{q, r}
	}

	q = newBigInt(false)
	r = x

	// Normalize Y.
	t := ceil(tb, bitsPerDigit) - 1
	lambda := 0
	for y.digits[t] < biHalfRadix {
		y = biShiftLeft(y, 1)
		lambda++
		tb++
		t = ceil(tb, bitsPerDigit) - 1
	}

	// Shift r over to keep the quotient constant. We'll shift the
	// remainder back at the end.
	r = biShiftLeft(r, lambda)
	nb += lambda // Update the bit count for x.
	n := ceil(nb, bitsPerDigit) - 1

	b := biMultiplyByRadixPower(y, n-t)
	for biCompare(r, b) != -1 {
		q.digits[n-t]++
		r = biSubtract(r, b)
	}

	i := n
	for i = n; i > t; i-- {
		var ri, ri1, ri2, yt, yt1 int
		if !(i >= len(r.digits)) {
			ri = r.digits[i]
		}
		if !(i-1 >= len(r.digits)) {
			ri1 = r.digits[i-1]
		}
		if !(i-2 >= len(r.digits)) {
			ri2 = r.digits[i-2]
		}
		if !(t >= len(y.digits)) {
			yt = y.digits[t]
		}
		if !(t-1 >= len(y.digits)) {
			yt1 = y.digits[t-1]
		}
		if ri == yt {
			q.digits[i-t-1] = maxDigitVal
		} else {
			q.digits[i-t-1] = (ri*biRadix + ri1) / yt
		}

		c1 := q.digits[i-t-1] * ((yt * biRadix) + yt1)
		c2 := (ri * biRadixSquared) + ((ri1 * biRadix) + ri2)
		for c1 > c2 {
			q.digits[i-t-1]--
			c1 = q.digits[i-t-1] * ((yt * biRadix) | yt1)
			c2 = (ri * biRadix * biRadix) + ((ri1 * biRadix) + ri2)
		}

		b = biMultiplyByRadixPower(y, i-t-1)
		r = biSubtract(r, biMultiplyDigit(b, q.digits[i-t-1]))
		if r.isNeg {
			r = biAdd(r, b)
			q.digits[i-t-1]--
		}
	}

	r = biShiftRight(r, lambda)
	// Fiddle with the signs and stuff to make sure that 0 <= r < y.
	q.isNeg = x.isNeg != origYIsNeg
	if x.isNeg {
		if origYIsNeg {
			q = biAdd(q, bigOne)
		} else {
			q = biSubtract(q, bigOne)
		}
		y = biShiftRight(y, lambda)
		r = biSubtract(y, r)
	}
	// Check for the unbelievably stupid degenerate case of r == -0.
	if r.digits[0] == 0 && biHighIndex(r) == 0 {
		r.isNeg = false
	}

	return [2]bigInt{q, r}
}

func biDivide(x bigInt, y bigInt) bigInt {
	return biDivideModulo(x, y)[0]
}

func biModulo(x bigInt, y bigInt) bigInt {
	return biDivideModulo(x, y)[1]
}

// func biMultiplyMod(x bigInt, y bigInt, m bigInt) bigInt {}

// func biPow(x BigInt, y BigInt) BigInt {}

// func biPowMod(x BigInt, y BigInt, m BigInt) BigInt {}
