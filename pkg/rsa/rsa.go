package rsa

import "math/rand"

var RSAAPP = map[string]int{
	"NoPadding":       1,
	"PKCS1Padding":    2,
	"RawEncoding":     1,
	"NumericEncoding": 2,
}

// KeyPair is invoked as the first step in the encryption or decryption
// process to take the three numbers (expressed as hexadecimal strings) that
// are used for RSA asymmetric encryption/decryption and turn them into a key
// object that can be used for encrypting and decrypting.
type KeyPair struct {
	encryptionExponent string
	decryptionExponent string
	modulus            string
	keylen             int
}

// E returns the encryptionExponent in bigInt format.
func (key *KeyPair) E() bigInt {
	return biFromHex(key.encryptionExponent)
}

// D returns the decryptionExponent in bigInt format.
func (key *KeyPair) D() bigInt {
	return biFromHex(key.decryptionExponent)
}

// M returns the modulis in bigInt format.
func (key *KeyPair) M() bigInt {
	return biFromHex(key.modulus)
}

// ChunkSize returns the current chunk size.
func (key *KeyPair) ChunkSize() int {
	// i := key.keylen / 8
	return 2 * biHighIndex(key.M())
}

// Radix returns biRadixBits.
func (key *KeyPair) Radix() int {
	return biRadixBits
}

// Barrett returns a new barretMu.
func (key *KeyPair) Barrett() barrettMu {
	return newBarretMu(key.M())
}

// NewKeyPair initializes a new RSAKeyPair.
func NewKeyPair(encryptionExponent string, decryptionExponent string, modulus string, keylen int) KeyPair {
	var key KeyPair
	key = KeyPair{
		encryptionExponent: encryptionExponent,
		decryptionExponent: decryptionExponent,
		modulus:            modulus,
		keylen:             keylen,
	}
	return key
}

/*****************************************************************************/

// EncryptedString accepts a plaintext string that is to be encrypted with the
// public key component of the previously-built RSA key using the RSA
// assymmetric encryption method.  Before it is encrypted, the plaintext
// string is padded to the same length as the encryption key for proper
// encryption.
func EncryptedString(key KeyPair, s string, pad int, encoding int) string {
	var result string                                  // Cypthertext result
	var i, j, k int                                    // The usual Fortran index stuff
	sl := len(s)                                       // Plaintext string length
	a := make([]int, key.ChunkSize(), key.ChunkSize()) // The usual Alice and Bob stuff

	// Figure out the padding type.
	padtype := 0
	if pad == RSAAPP["NoPadding"] {
		padtype = 1
	} else if pad == RSAAPP["PKCS1Padding"] {
		padtype = 2
	}

	// Determine encoding type.
	encodingtype := 0
	if encoding == RSAAPP["RawEncoding"] {
		encodingtype = 1
	}

	// If we're not using Dave's padding method, we need to truncate long
	// plaintext blocks to the correct length for the padding method used:
	if padtype == 1 {
		if sl > key.ChunkSize() {
			sl = key.ChunkSize()
		}
	} else if padtype == 2 {
		if sl > key.ChunkSize()-11 {
			sl = key.ChunkSize() - 11
		}
	}

	// * Convert the plaintext string to an array of characters so that we can work
	// * with individual characters.
	if padtype == 2 {
		j = sl - 1
	} else {
		j = key.ChunkSize() - 1
	}

	for i < sl {
		if padtype > 0 {
			a[j] = int(s[i])
		} else {
			a[i] = int(s[i])
		}
		i++
		j--
	}

	// Now is the time to add the padding.
	// The padding is either a zero byte or a randomly-generated non-zero byte.

	if padtype == 1 {
		i = 0
	}

	j = key.ChunkSize() - (sl % key.ChunkSize())

	for j > 0 {
		if padtype == 2 {
			rpad := rand.Intn(256)

			a[i] = rpad
		} else {
			a[i] = 0
		}

		i++
		j--
	}

	// For PKCS1v1.5 padding, we need to fill in the block header.
	//
	// According to RFC 2313, a block type, a padding string, and the data shall
	// be formatted into the encryption block:
	//
	//      EncrBlock = 00 || BlockType || PadString || 00 || Data
	if padtype == 2 {
		a[sl] = 0
		a[key.ChunkSize()-2] = 2
		a[key.ChunkSize()-1] = 0
	}

	// Initialize Barrett here (different from JavaScript)
	barrett := key.Barrett()

	// Carve up the plaintext and encrypt each of the resultant blocks.
	al := len(a)

	for i = 0; i < al; i += key.ChunkSize() {
		// Get a block.
		block := newBigInt(false)

		j = 0

		for k = i; k < i+key.ChunkSize(); j++ {
			block.digits[j] = a[k]
			k++
			block.digits[j] += a[k] << 8
			k++
		}

		var text string
		// Encrypt it, convert it to text, and append it to the result.
		crypt := barrett.powMod(block, key.E())
		if encodingtype == 1 {
			text = biToBytes(crypt)
		} else {
			if key.Radix() == 16 {
				text = biToHex(crypt)
			} else {
				text = biToString(crypt, key.Radix())
			}
		}
		result += text
	}

	return result
}
