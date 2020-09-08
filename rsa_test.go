package rsa

import (
	"reflect"
	"testing"

	"github.com/stellirin/RSA.go/barrett"
	"github.com/stellirin/RSA.go/bigint"
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
		want   bigint.BigInt
	}

	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: bigint.BiFromHex(encryptionExponent),
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
		want   bigint.BigInt
	}

	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: bigint.BiFromHex(decryptionExponent),
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
		want   bigint.BigInt
	}

	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: bigint.BiFromHex(modulus),
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
		want   barrett.BarrettMu
	}

	tests := []test{
		{
			fields: fields{
				encryptionExponent: encryptionExponent,
				decryptionExponent: decryptionExponent,
				modulus:            modulus,
				keylen:             keylen,
			},
			want: barrett.NewBarretMu(bigint.BiFromHex(modulus)),
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

	key := KeyPair{
		encryptionExponent: encryptionExponent,
		decryptionExponent: decryptionExponent,
		modulus:            modulus,
		keylen:             keylen,
	}
	tests := []test{
		{
			name: "00",
			args: args{
				key:      key,
				s:        "password",
				pad:      0,
				encoding: 0,
			},
			want: "1afa376394f638f4a2828c31abadf05ba84b8a7698a1f8c7374c172912df10f4b821dfca1d9830f53e87d8311b4af6f8a07c3e46721eb1517f100ce8f7fac62e6a2d32a210929efb01275884ef5a2284f269eeb9380c15bcdfc52a49ea04429849059166394ee1f220d8e92a64583646d0499fcd0b345e474c1f4d6074d4bcb8",
		},
		{
			name: "01",
			args: args{
				key:      key,
				s:        "password",
				pad:      0,
				encoding: 1,
			},
			want: "\x1aú7c\u0094ö8ô¢\u0082\u008c1«\u00adð[¨K\u008av\u0098¡øÇ7L\x17)\x12ß\x10ô¸!ßÊ\x1d\u00980õ>\u0087Ø1\x1bJöø\u00a0|>Fr\x1e±Q\u007f\x10\fè÷úÆ.j-2¢\x10\u0092\u009eû\x01'X\u0084ïZ\"\u0084òiî¹8\f\x15¼ßÅ*Iê\x04B\u0098I\x05\u0091f9Náò Øé*dX6FÐI\u009fÍ\v4^GL\x1fM`tÔ¼¸",
		},
		{
			name: "10",
			args: args{
				key:      key,
				s:        "password",
				pad:      1,
				encoding: 0,
			},
			want: "38aa05dc22a43d185bec003f7fb9d2858e3a24a52aae23b152f73f39ce38e7d635b549ba35658ee11ecab0e55ca07b0f9b9ce005135db8316fa57c802e81de8753c981e8a4ef9ce30714202e72192616584d175cd553961e69fe2f3a80bfded9d57ba551c29a1ee28beae7e5d303f991e75597561cb3c3da7e4686429bdf5480",
		},
		{
			name: "11",
			args: args{
				key:      key,
				s:        "password",
				pad:      1,
				encoding: 1,
			},
			want: "8ª\x05Ü\"¤=\x18[ì\x00?\u007f¹Ò\u0085\u008e:$¥*®#±R÷?9Î8çÖ5µIº5e\u008eá\x1eÊ°å\\\u00a0{\x0f\u009b\u009cà\x05\x13]¸1o¥|\u0080.\u0081Þ\u0087SÉ\u0081è¤ï\u009cã\a\x14 .r\x19&\x16XM\x17\\ÕS\u0096\x1eiþ/:\u0080¿ÞÙÕ{¥QÂ\u009a\x1eâ\u008bêçåÓ\x03ù\u0091çU\u0097V\x1c³ÃÚ~F\u0086B\u009bßT\u0080",
			// want: "",
		},
		{
			name: "20",
			args: args{
				key:      key,
				s:        "password",
				pad:      2,
				encoding: 0,
			},
			want: "14a071f639b866b2e4a1c4a9a26e3783a9d624f58fed329a749836612ab0095be7f2434c9da1077606872cd414a0cfffe222744454377f2a7bbcaa3fb4fe20e30bf567bf2b055974ba653eec95a1dbae5e5c91963953de55b4cfb26e66752a2045c4e4b769ea028a31f934ab741545e412b9b028357dab26433142ec1a9a3232",
		},
		{
			name: "21",
			args: args{
				key:      key,
				s:        "password",
				pad:      2,
				encoding: 1,
			},
			want: "\x16\\\u0083\x10M»7æ#±\u0083\"8à\x1f\u0088ÃÍV\u0097$?e/\aù\u0084Ó\u009e\u0084&5\x11\u0095\x12\u0086Ñ$Ï²H$¼\x1clÚelµ|Bx*¢í<u\u0093m\u0093\x17\x12kÇve\nÚ[F\nex*ò\u0082íÅöã\x1bLUï\u0081b¾E¯\u00811Ã2Ç\u00818[\n\x1cÐüX ÖY¶$ù\u0097ý¬11OÛQc&Y4:Òò\u009dÿÍvo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptedString(tt.args.key, tt.args.s, tt.args.pad, tt.args.encoding); got != tt.want {
				t.Errorf("EncryptedString() = %q, want %v", got, tt.want)
			}
		})
	}
}
