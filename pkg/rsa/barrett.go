package rsa

// BarrettMu is a struct.
type BarrettMu struct {
	modulus BigInt
	k       int
	mu      BigInt
	bkplus1 BigInt
}

func (b *BarrettMu) modulo(x BigInt) BigInt {
	q1 := biDivideByRadixPower(x, b.k-1)
	q2 := biMultiply(q1, b.mu)
	q3 := biDivideByRadixPower(q2, b.k+1)
	r1 := biModuloByRadixPower(x, b.k+1)
	r2term := biMultiply(q3, b.modulus)
	r2 := biModuloByRadixPower(r2term, b.k+1)
	r := biSubtract(r1, r2)
	if r.isNeg {
		r = biAdd(r, b.bkplus1)
	}
	rgtem := biCompare(r, b.modulus) >= 0
	for rgtem {
		r = biSubtract(r, b.modulus)
		rgtem = biCompare(r, b.modulus) >= 0
	}
	return r
}

func (b *BarrettMu) multiplyMod(x BigInt, y BigInt) BigInt {
	xy := biMultiply(x, y)
	return b.modulo(xy)
}

func (b *BarrettMu) powMod(x BigInt, y BigInt) BigInt {
	result := NewBigInt(false)
	result.digits[0] = 1
	a := x
	k := y
	for {
		if (k.digits[0] & 1) != 0 {
			result = b.multiplyMod(result, a)
		}
		k = biShiftRight(k, 1)
		if k.digits[0] == 0 && biHighIndex(k) == 0 {
			break
		}
		a = b.multiplyMod(a, a)
	}
	return result
}

func newBarretMu(m BigInt) BarrettMu {
	modulus := biCopy(m)
	k := biHighIndex(modulus) + 1
	b2k := NewBigInt(false)
	b2k.digits[2*k] = 1 // b2k = b^(2k)
	mu := biDivide(b2k, modulus)
	bkplus1 := NewBigInt(false)
	bkplus1.digits[k+1] = 1

	b := BarrettMu{
		modulus: modulus,
		k:       k,
		mu:      mu,
		bkplus1: bkplus1,
	}
	return b
}
