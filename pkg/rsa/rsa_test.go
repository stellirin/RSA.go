package rsa

import (
	"testing"
)

var encryptionExponent = "10001"
var decryptionExponent = "0"
var modulus = "b1ae618873ee7ff972e9be6dba93f24f0ef38ac50c2265f7274696b37e6159d2c81a798552360e941be4f9e22522b5421bf753f1ab22626ddee300ee675553e57ad5ab86f77a75c28babdb3e263aad1245e4a2cf58789406083d56d3a4bd6d04e17a77f1919a2b9e1176a0b21931bc82a132ef0af661076d92cef6b13313886f"

func TestKeyPair(t *testing.T) {
	var expected []int
	var result []int

	SetMaxDigits(262)
	key := NewKeyPair("deadbeef", "beefdead", modulus, 0)

	expected = []int{48879, 57005} // beef, dead
	result = key.E().digits[0:len(expected)]
	for i, e := range expected {
		if e != result[i] {
			t.Errorf("KeyPair().E() failed, expected %v, got %v", expected, result)
			break
		}
	}

	expected = []int{57005, 48879} // dead, beef
	result = key.D().digits[0:len(expected)]
	for i, e := range expected {
		if e != result[i] {
			t.Errorf("KeyPair().D() failed, expected %v, got %v", expected, result)
			break
		}
	}

	expected = []int{34927, 13075, 63153, 37582, 1901, 63073, 61194, 41266, 48258, 6449, 41138, 4470, 11166, 37274, 30705, 57722, 27908, 42173, 22227, 2109, 37894, 22648, 41679, 17892, 44306, 9786, 56126, 35755, 30146, 63354, 43910, 31445, 21477, 26453, 238, 57059, 25197, 43810, 21489, 7159, 46402, 9506, 63970, 7140, 3732, 21046, 31109, 51226, 22994, 32353, 38579, 10054, 26103, 3106, 35525, 3827, 62031, 47763, 48749, 29417, 32761, 29678, 24968, 45486}
	result = key.M().digits[0:len(expected)]
	for i, e := range expected {
		if e != result[i] {
			t.Errorf("KeyPair().M() failed, expected %v, got %v", expected, result)
			break
		}
	}

	expected = []int{126}
	result = []int{key.ChunkSize()}
	for i, e := range expected {
		if e != result[i] {
			t.Errorf("KeyPair().ChunkSize() failed, expected %v, got %v", expected, result)
			break
		}
	}

	expected = []int{16}
	result = []int{key.Radix()}
	for i, e := range expected {
		if e != result[i] {
			t.Errorf("KeyPair().Radix() failed, expected %v, got %v", expected, result)
			break
		}
	}
}

func TestEncryptedString(t *testing.T) {
	SetMaxDigits(262)
	s := "password"
	key := NewKeyPair(encryptionExponent, decryptionExponent, modulus, 0)

	expected := "1afa376394f638f4a2828c31abadf05ba84b8a7698a1f8c7374c172912df10f4b821dfca1d9830f53e87d8311b4af6f8a07c3e46721eb1517f100ce8f7fac62e6a2d32a210929efb01275884ef5a2284f269eeb9380c15bcdfc52a49ea04429849059166394ee1f220d8e92a64583646d0499fcd0b345e474c1f4d6074d4bcb8"
	result := EncryptedString(key, s, 0, 0)

	if expected != result {
		t.Errorf("EncryptedString failed, expected %v, got %v", expected, result)
		// break
	}
}
