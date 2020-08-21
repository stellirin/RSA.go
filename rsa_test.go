package rsa

import (
	"reflect"
	"testing"

	. "github.com/stellirin/RSA.go/barrett"
	. "github.com/stellirin/RSA.go/bigint"
)

var encryptionExponent = "10001"
var decryptionExponent = "0"
var modulus = "b1ae618873ee7ff972e9be6dba93f24f0ef38ac50c2265f7274696b37e6159d2c81a798552360e941be4f9e22522b5421bf753f1ab22626ddee300ee675553e57ad5ab86f77a75c28babdb3e263aad1245e4a2cf58789406083d56d3a4bd6d04e17a77f1919a2b9e1176a0b21931bc82a132ef0af661076d92cef6b13313886f"
var keylen = 0

func Test_KeyPair_E(t *testing.T) {
	type fields struct {
		encryptionExponent string
		decryptionExponent string
		modulus            string
		keylen             int
	}
	type test struct {
		name   string
		fields fields
		want   BigInt
	}

	SetMaxDigits(256)
	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: BiFromHex(encryptionExponent),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &KeyPair{
				encryptionExponent: tt.fields.encryptionExponent,
				decryptionExponent: tt.fields.decryptionExponent,
				modulus:            tt.fields.modulus,
				keylen:             tt.fields.keylen,
			}
			if got := key.E(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyPair.E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_KeyPair_D(t *testing.T) {
	type fields struct {
		encryptionExponent string
		decryptionExponent string
		modulus            string
		keylen             int
	}
	type test struct {
		name   string
		fields fields
		want   BigInt
	}

	SetMaxDigits(256)
	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: BiFromHex(decryptionExponent),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &KeyPair{
				encryptionExponent: tt.fields.encryptionExponent,
				decryptionExponent: tt.fields.decryptionExponent,
				modulus:            tt.fields.modulus,
				keylen:             tt.fields.keylen,
			}
			if got := key.D(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyPair.D() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_KeyPair_M(t *testing.T) {
	type fields struct {
		encryptionExponent string
		decryptionExponent string
		modulus            string
		keylen             int
	}
	type test struct {
		name   string
		fields fields
		want   BigInt
	}

	SetMaxDigits(256)
	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: BiFromHex(modulus),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &KeyPair{
				encryptionExponent: tt.fields.encryptionExponent,
				decryptionExponent: tt.fields.decryptionExponent,
				modulus:            tt.fields.modulus,
				keylen:             tt.fields.keylen,
			}
			if got := key.M(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyPair.M() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_KeyPair_ChunkSize(t *testing.T) {
	type fields struct {
		encryptionExponent string
		decryptionExponent string
		modulus            string
		keylen             int
	}
	type test struct {
		name   string
		fields fields
		want   int
	}

	SetMaxDigits(256)
	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: 126,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &KeyPair{
				encryptionExponent: tt.fields.encryptionExponent,
				decryptionExponent: tt.fields.decryptionExponent,
				modulus:            tt.fields.modulus,
				keylen:             tt.fields.keylen,
			}
			if got := key.ChunkSize(); got != tt.want {
				t.Errorf("KeyPair.ChunkSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_KeyPair_Radix(t *testing.T) {
	type fields struct {
		encryptionExponent string
		decryptionExponent string
		modulus            string
		keylen             int
	}
	type test struct {
		name   string
		fields fields
		want   int
	}

	SetMaxDigits(256)
	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &KeyPair{
				encryptionExponent: tt.fields.encryptionExponent,
				decryptionExponent: tt.fields.decryptionExponent,
				modulus:            tt.fields.modulus,
				keylen:             tt.fields.keylen,
			}
			if got := key.Radix(); got != tt.want {
				t.Errorf("KeyPair.Radix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_KeyPair_Barrett(t *testing.T) {
	type fields struct {
		encryptionExponent string
		decryptionExponent string
		modulus            string
		keylen             int
	}
	type test struct {
		name   string
		fields fields
		want   BarrettMu
	}

	SetMaxDigits(256)
	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: NewBarretMu(BiFromHex(modulus)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &KeyPair{
				encryptionExponent: tt.fields.encryptionExponent,
				decryptionExponent: tt.fields.decryptionExponent,
				modulus:            tt.fields.modulus,
				keylen:             tt.fields.keylen,
			}
			if got := key.Barrett(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyPair.Barrett() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewKeyPair(t *testing.T) {
	type args struct {
		encryptionExponent string
		decryptionExponent string
		modulus            string
		keylen             int
	}
	type test struct {
		name string
		args args
		want KeyPair
	}

	SetMaxDigits(256)
	tests := []test{
		{
			args: args{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: KeyPair{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeyPair(tt.args.encryptionExponent, tt.args.decryptionExponent, tt.args.modulus, tt.args.keylen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeyPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_EncryptedString(t *testing.T) {
	type args struct {
		key      KeyPair
		s        string
		pad      int
		encoding int
	}
	type test struct {
		name string
		args args
		want string
	}

	SetMaxDigits(256)
	tests := []test{
		{
			args: args{
				key: KeyPair{
					encryptionExponent: encryptionExponent,
					decryptionExponent: decryptionExponent,
					modulus:            modulus,
					keylen:             keylen,
				},
				s:        "password",
				pad:      0,
				encoding: 0,
			},
			want: "1afa376394f638f4a2828c31abadf05ba84b8a7698a1f8c7374c172912df10f4b821dfca1d9830f53e87d8311b4af6f8a07c3e46721eb1517f100ce8f7fac62e6a2d32a210929efb01275884ef5a2284f269eeb9380c15bcdfc52a49ea04429849059166394ee1f220d8e92a64583646d0499fcd0b345e474c1f4d6074d4bcb8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptedString(tt.args.key, tt.args.s, tt.args.pad, tt.args.encoding); got != tt.want {
				t.Errorf("EncryptedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
