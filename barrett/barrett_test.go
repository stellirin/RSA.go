package barrett

import (
	"reflect"
	"testing"

	"github.com/stellirin/RSA.go/bigint"
)

func Test_BarrettMu_Modulo(t *testing.T) {
	type fields struct {
		modulus bigint.BigInt
		k       int
		mu      bigint.BigInt
		bkplus1 bigint.BigInt
	}
	type args struct {
		x bigint.BigInt
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bigint.BigInt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BarrettMu{
				modulus: tt.fields.modulus,
				k:       tt.fields.k,
				mu:      tt.fields.mu,
				bkplus1: tt.fields.bkplus1,
			}
			if got := b.Modulo(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BarrettMu.Modulo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BarrettMu_MultiplyMod(t *testing.T) {
	type fields struct {
		modulus bigint.BigInt
		k       int
		mu      bigint.BigInt
		bkplus1 bigint.BigInt
	}
	type args struct {
		x bigint.BigInt
		y bigint.BigInt
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bigint.BigInt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BarrettMu{
				modulus: tt.fields.modulus,
				k:       tt.fields.k,
				mu:      tt.fields.mu,
				bkplus1: tt.fields.bkplus1,
			}
			if got := b.MultiplyMod(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BarrettMu.MultiplyMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BarrettMu_PowMod(t *testing.T) {
	type fields struct {
		modulus bigint.BigInt
		k       int
		mu      bigint.BigInt
		bkplus1 bigint.BigInt
	}
	type args struct {
		x bigint.BigInt
		y bigint.BigInt
	}
	type test struct {
		name   string
		fields fields
		args   args
		want   bigint.BigInt
	}

	bigint.SetMaxDigits(256)
	tests := []test{
		{
			fields: fields{
				modulus: bigint.BiFromHex("b1ae618873ee7ff972e9be6dba93f24f0ef38ac50c2265f7274696b37e6159d2c81a798552360e941be4f9e22522b5421bf753f1ab22626ddee300ee675553e57ad5ab86f77a75c28babdb3e263aad1245e4a2cf58789406083d56d3a4bd6d04e17a77f1919a2b9e1176a0b21931bc82a132ef0af661076d92cef6b13313886f"),
				k:       64,
				mu:      bigint.BiFromHex("000170d725f48c65ed6b3ca1d61f5e0b793b11d10ca10aafdefe1052e1c2bac4f5b2f94e8088ab02e78a4e133c88245c5e78d46c23eb2e9bcf8eac7bf793857ecd718bc7e0c975031cf1921465051aa1db4bb53dc049ef41f8c98ab619ebeb8782c9842a5a91c329038bcc0f03d4ab7282f6b7c699a00a5883e9f31478720b9c0416"),
				bkplus1: bigint.BiFromHex("1"),
			},
			args: args{
				x: bigint.BiFromHex("16"),
				y: bigint.BiFromHex("10001"),
			},
			want: bigint.BiFromHex("991d42f6374d8bd58f68b6e908d986edf82b722e0d8ec17d5ab383998620a4d184ff804b713bb8e1be38b689308cc9ceac0e85924501eaddf5ef3dfd2791eea7f1ca3d79fe0315cb73359e0e788e178895bf258fc2fd57563e6c72cc7ac43c59d9b98a8bf31b6041c69850afc94e9186db80c5f894e114b7de875270907badfe"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BarrettMu{
				modulus: tt.fields.modulus,
				k:       tt.fields.k,
				mu:      tt.fields.mu,
				bkplus1: tt.fields.bkplus1,
			}
			if got := b.PowMod(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BarrettMu.PowMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_New(t *testing.T) {
	type args struct {
		m bigint.BigInt
	}
	type test struct {
		name string
		args args
		want BarrettMu
	}

	tests := []test{
		{
			args: args{
				m: bigint.BiFromHex("b1ae618873ee7ff972e9be6dba93f24f0ef38ac50c2265f7274696b37e6159d2c81a798552360e941be4f9e22522b5421bf753f1ab22626ddee300ee675553e57ad5ab86f77a75c28babdb3e263aad1245e4a2cf58789406083d56d3a4bd6d04e17a77f1919a2b9e1176a0b21931bc82a132ef0af661076d92cef6b13313886f"),
			},
			want: BarrettMu{
				modulus: bigint.BiFromHex("b1ae618873ee7ff972e9be6dba93f24f0ef38ac50c2265f7274696b37e6159d2c81a798552360e941be4f9e22522b5421bf753f1ab22626ddee300ee675553e57ad5ab86f77a75c28babdb3e263aad1245e4a2cf58789406083d56d3a4bd6d04e17a77f1919a2b9e1176a0b21931bc82a132ef0af661076d92cef6b13313886f"),
				k:       64,
				mu:      bigint.BiFromHex("000170d725f48c65ed6b3ca1d61f5e0b793b11d10ca10aafdefe1052e1c2bac4f5b2f94e8088ab02e78a4e133c88245c5e78d46c23eb2e9bcf8eac7bf793857ecd718bc7e0c975031cf1921465051aa1db4bb53dc049ef41f8c98ab619ebeb8782c9842a5a91c329038bcc0f03d4ab7282f6b7c699a00a5883e9f31478720b9c0416"),
				bkplus1: bigint.BiMultiplyByRadixPower(bigint.BiFromHex("1"), 64+1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBarretMu() = %v, want %v", got, tt.want)
			}
		})
	}
}
