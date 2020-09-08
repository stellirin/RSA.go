package barrett

import (
	"github.com/stellirin/RSA.go/bigint"
)

// BarrettMu is a struct.
type BarrettMu struct {
	modulus bigint.BigInt
	k       int
	mu      bigint.BigInt
	bkplus1 bigint.BigInt
}

// Modulo calculates the modulo of a BarretMu.
func (b *BarrettMu) Modulo(x bigint.BigInt) bigint.BigInt {
	q1 := bigint.BiDivideByRadixPower(x, b.k-1)
	q2 := bigint.BiMultiply(q1, b.mu)
	q3 := bigint.BiDivideByRadixPower(q2, b.k+1)
	r1 := bigint.BiModuloByRadixPower(x, b.k+1)
	r2term := bigint.BiMultiply(q3, b.modulus)
	r2 := bigint.BiModuloByRadixPower(r2term, b.k+1)
	r := bigint.BiSubtract(r1, r2)
	if r.IsNeg {
		r = bigint.BiAdd(r, b.bkplus1)
	}
	rgtem := bigint.BiCompare(r, b.modulus) >= 0
	for rgtem {
		r = bigint.BiSubtract(r, b.modulus)
		rgtem = bigint.BiCompare(r, b.modulus) >= 0
	}
	return r
}

// MultiplyMod multiplies the mod.
func (b *BarrettMu) MultiplyMod(x bigint.BigInt, y bigint.BigInt) bigint.BigInt {
	xy := bigint.BiMultiply(x, y)
	return b.Modulo(xy)
}

// PowMod powers the mod.
func (b *BarrettMu) PowMod(x bigint.BigInt, y bigint.BigInt) bigint.BigInt {
	result := bigint.New(false)
	result.Digits[0] = 1
	a := x
	k := y
	for {
		if (k.Digits[0] & 1) != 0 {
			result = b.MultiplyMod(result, a)
		}
		k = bigint.BiShiftRight(k, 1)
		if k.Digits[0] == 0 && bigint.BiHighIndex(k) == 0 {
			break
		}
		a = b.MultiplyMod(a, a)
	}
	return result
}

// New initializes a new BarretMu.
func New(m bigint.BigInt) BarrettMu {
	modulus := bigint.BiCopy(m)
	k := bigint.BiHighIndex(modulus) + 1
	b2k := bigint.New(false)
	b2k.Digits[2*k] = 1 // b2k = b^(2k)
	mu := bigint.BiDivide(b2k, modulus)
	bkplus1 := bigint.New(false)
	bkplus1.Digits[k+1] = 1

	b := BarrettMu{
		modulus: modulus,
		k:       k,
		mu:      mu,
		bkplus1: bkplus1,
	}
	return b
}
