package rsa

// BarrettMu is a struct.
type BarrettMu struct {
	modulus BigInt
	k       int
	mu      BigInt
	bkplus1 BigInt
}

func (b *BarrettMu) Modulo(x BigInt) BigInt {
	q1 := BiDivideByRadixPower(x, b.k-1)
	q2 := BiMultiply(q1, b.mu)
	q3 := BiDivideByRadixPower(q2, b.k+1)
	r1 := BiModuloByRadixPower(x, b.k+1)
	r2term := BiMultiply(q3, b.modulus)
	r2 := BiModuloByRadixPower(r2term, b.k+1)
	r := BiSubtract(r1, r2)
	if r.IsNeg {
		r = BiAdd(r, b.bkplus1)
	}
	rgtem := BiCompare(r, b.modulus) >= 0
	for rgtem {
		r = BiSubtract(r, b.modulus)
		rgtem = BiCompare(r, b.modulus) >= 0
	}
	return r
}

func (b *BarrettMu) MultiplyMod(x BigInt, y BigInt) BigInt {
	xy := BiMultiply(x, y)
	return b.Modulo(xy)
}

func (b *BarrettMu) PowMod(x BigInt, y BigInt) BigInt {
	result := NewBigInt(false)
	result.Digits[0] = 1
	a := x
	k := y
	for {
		if (k.Digits[0] & 1) != 0 {
			result = b.MultiplyMod(result, a)
		}
		k = BiShiftRight(k, 1)
		if k.Digits[0] == 0 && BiHighIndex(k) == 0 {
			break
		}
		a = b.MultiplyMod(a, a)
	}
	return result
}

func NewBarretMu(m BigInt) BarrettMu {
	modulus := BiCopy(m)
	k := BiHighIndex(modulus) + 1
	b2k := NewBigInt(false)
	b2k.Digits[2*k] = 1 // b2k = b^(2k)
	mu := BiDivide(b2k, modulus)
	bkplus1 := NewBigInt(false)
	bkplus1.Digits[k+1] = 1

	b := BarrettMu{
		modulus: modulus,
		k:       k,
		mu:      mu,
		bkplus1: bkplus1,
	}
	return b
}
