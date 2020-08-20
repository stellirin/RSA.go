package rsa

import (
	"reflect"
	"testing"
)

// Used to store benchmark results
var rBigInt BigInt
var rString string
var rInt int

var testRunes = []rune{222, 173, 190, 239, 202, 254, 240, 13, 186, 190, 192, 222, 250, 206, 190, 173}
var testString = "deadbeefcafef00dbabec0defacebead"
var testBigInt = BigInt{
	Digits: []int{48813, 64206, 49374, 47806, 61453, 51966, 48879, 57005}, // bead, face, c0de, babe, f00d, cafe, beef, dead
	IsNeg:  false,
}

// func Test_biFromDecimal(t *testing.T) {}

func Benchmark_biCopy(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiCopy(testBigInt)
	}
}

func Test_biCopy(t *testing.T) {
	type args struct {
		bi BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				bi: testBigInt,
			},
			want: testBigInt,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiCopy(tt.args.bi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biFromNumber(t *testing.T) {}

// func Benchmark_biToString(b *testing.B) {}

func Test_biToString(t *testing.T) {
	type args struct {
		x     BigInt
		radix int
	}
	type test struct {
		name string
		args args
		want string
	}
	SetMaxDigits(8)

	tests := []test{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiToString(tt.args.x, tt.args.radix); got != tt.want {
				t.Errorf("biToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biToDecimal(t *testing.T) {}

func Benchmark_biToHex(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rString = BiToHex(testBigInt)
	}
}

func Test_biToHex(t *testing.T) {
	type args struct {
		x BigInt
	}
	type test struct {
		name string
		args args
		want string
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
			},
			want: testString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiToHex(tt.args.x); got != tt.want {
				t.Errorf("biToHex() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Benchmark_biFromHex(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiFromHex(testString)
	}
}

func Test_biFromHex(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				s: testString,
			},
			want: testBigInt,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiFromHex(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biFromHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biFromString(t *testing.T) {}

func Benchmark_biToBytes(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rString = BiToBytes(testBigInt)
	}
}

func Test_biToBytes(t *testing.T) {
	type args struct {
		x BigInt
	}
	type test struct {
		name string
		args args
		want string
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
			},
			want: string(testRunes),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiToBytes(tt.args.x); got != tt.want {
				t.Errorf("biToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biDump(t *testing.T) {}

func Benchmark_biAdd(b *testing.B) {
	SetMaxDigits(16)
	x := BiFromHex("deadbeef")
	y := BiFromHex("beefdead")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiAdd(x, y)
	}
}

func Test_biAdd(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(16)

	tests := []test{
		{
			name: "Pos+Pos",
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex(testString),
			},
			want: BigInt{
				Digits: []int{32090, 62877, 33213, 30077, 57371, 38397, 32223, 48475, 1, 0, 0, 0, 0, 0, 0, 0},
				IsNeg:  false,
			},
		},
		{
			name: "Pos+Neg",
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex("-" + reverseStr(testString)),
			},
			want: BigInt{
				Digits: []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961, 0, 0, 0, 0, 0, 0, 0, 0},
				IsNeg:  false,
			},
		},
		{
			name: "Neg+Neg",
			args: args{
				x: BiFromHex("-" + reverseStr(testString)),
				y: BiFromHex("-" + reverseStr(testString)),
			},
			want: BigInt{
				Digits: []int{46554, 64983, 57177, 40991, 55127, 55833, 55647, 46551, 1, 0, 0, 0, 0, 0, 0, 0},
				IsNeg:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiAdd(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biSubtract(b *testing.B) {
	SetMaxDigits(16)
	x := BiFromHex(testString)
	y := BiFromHex(reverseStr(testString))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiSubtract(x, y)
	}
}

func Test_biSubtract(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(16)

	tests := []test{
		{
			name: "Pos-Pos",
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex(reverseStr(testString)),
			},
			want: BigInt{
				Digits: []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961, 0, 0, 0, 0, 0, 0, 0, 0},
				IsNeg:  false,
			},
		},
		{
			name: "Pos-Neg",
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex("-" + reverseStr(testString)),
			},
			want: BigInt{
				Digits: []int{39322, 63930, 45195, 35534, 56249, 47115, 43935, 47513, 1, 0, 0, 0, 0, 0, 0, 0},
				IsNeg:  false,
			},
		},
		{
			name: "Neg-Neg",
			args: args{
				x: BiFromHex("-" + reverseStr(testString)),
				y: BiFromHex("-" + reverseStr(testString)),
			},
			want: BigInt{
				Digits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				IsNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiSubtract(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biSubtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biHighIndex(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = BiHighIndex(testBigInt)
	}
}

func Test_biHighIndex(t *testing.T) {
	type args struct {
		x BigInt
	}
	type test struct {
		name string
		args args
		want int
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
			},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiHighIndex(tt.args.x); got != tt.want {
				t.Errorf("biHighIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biNumBits(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = BiNumBits(testBigInt)
	}
}

func Test_biNumBits(t *testing.T) {
	type args struct {
		x BigInt
	}
	type test struct {
		name string
		args args
		want int
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
			},
			want: 128,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiNumBits(tt.args.x); got != tt.want {
				t.Errorf("biNumBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biMultiply(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiMultiply(testBigInt, testBigInt)
	}
}

func Test_biMultiply(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(16)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				y: testBigInt,
			},
			want: BigInt{
				Digits: []int{
					16617, 20593, 42861, 32538, 45046, 25231, 13697, 47718,
					1165, 9698, 10208, 22465, 22231, 33427, 52499, 49585,
				},
				IsNeg: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiMultiply(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biMultiplyDigit(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiMultiplyDigit(testBigInt, 16)
	}
}

func Test_biMultiplyDigit(t *testing.T) {
	type args struct {
		x BigInt
		y int
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(16)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				y: 16,
			},
			want: BigInt{
				Digits: []int{
					60112, 44267, 3567, 44012, 219, 45039, 61180, 60123,
					13, 0, 0, 0, 0, 0, 0, 0,
				},
				IsNeg: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiMultiplyDigit(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biMultiplyDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biShiftLeft(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiShiftLeft(testBigInt, 16)
	}
}

func Test_biShiftLeft(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 16,
			},
			want: BigInt{
				Digits: []int{0, 48813, 64206, 49374, 47806, 61453, 51966, 48879},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiShiftLeft(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biShiftLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biShiftRight(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiShiftRight(testBigInt, 16)
	}
}

func Test_biShiftRight(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 16,
			},
			want: BigInt{
				Digits: []int{64206, 49374, 47806, 61453, 51966, 48879, 57005, 0},
				IsNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiShiftRight(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biShiftRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biMultiplyByRadixPower(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiMultiplyByRadixPower(testBigInt, 3)
	}
}

func Test_biMultiplyByRadixPower(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 3,
			},
			want: BigInt{
				Digits: []int{0, 0, 0, 48813, 64206, 49374, 47806, 61453},
				IsNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiMultiplyByRadixPower(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biMultiplyByRadixPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biDivideByRadixPower(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiDivideByRadixPower(testBigInt, 3)
	}
}

func Test_biDivideByRadixPower(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 3,
			},
			want: BigInt{
				Digits: []int{47806, 61453, 51966, 48879, 57005, 0, 0, 0},
				IsNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiDivideByRadixPower(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biDivideByRadixPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biModuloByRadixPower(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiModuloByRadixPower(testBigInt, 3)
	}
}

func Test_biModuloByRadixPower(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 3,
			},
			want: BigInt{
				Digits: []int{48813, 64206, 49374, 0, 0, 0, 0, 0},
				IsNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiModuloByRadixPower(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biModuloByRadixPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biCompare(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = BiCompare(testBigInt, testBigInt)
	}
}

func Test_biCompare(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want int
	}
	SetMaxDigits(8)

	w, x, y, z := "-bead", "-face", "dead", "beef"
	biw, bix, biy, biz := BiFromHex(w), BiFromHex(x), BiFromHex(y), BiFromHex(z)

	tests := []test{
		{
			name: "-x < y",
			args: args{
				x: bix,
				y: biy,
			},
			want: -1,
		},
		{
			name: "x > -y",
			args: args{
				x: biz,
				y: bix,
			},
			want: 1,
		},
		{
			name: "-x < -y",
			args: args{
				x: bix,
				y: biw,
			},
			want: -1,
		},
		{
			name: "-x > -y",
			args: args{
				x: biw,
				y: bix,
			},
			want: 1,
		},
		{
			name: "x < y",
			args: args{
				x: biz,
				y: biy,
			},
			want: -1,
		},
		{
			name: "x > y",
			args: args{
				x: biy,
				y: biz,
			},
			want: 1,
		},
		{
			name: "x == y",
			args: args{
				x: biy,
				y: biy,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiCompare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("biCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biDivideModulo(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiDivideModulo(testBigInt, testBigInt)[0]
	}
}

func Test_biDivideModulo(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want [2]BigInt
	}
	SetMaxDigits(16)

	tests := []test{
		{
			name: "Pos/Pos",
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex(reverseStr(testString)),
			},
			want: [2]BigInt{
				bigOne,
				{
					Digits: []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961, 0, 0, 0, 0, 0, 0, 0, 0},
					IsNeg:  false,
				},
			},
		},
		{
			name: "Neg/Neg",
			args: args{
				x: BiFromHex("-" + testString),
				y: BiFromHex("-" + reverseStr(testString)),
			},
			want: [2]BigInt{
				{
					Digits: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					IsNeg:  false,
				},
				{
					Digits: []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961, 0, 0, 0, 0, 0, 0, 0, 0},
					IsNeg:  true,
				},
			},
		},
		{
			name: "Small/Big",
			args: args{
				x: BiFromHex(testString[0:24]),
				y: BiFromHex(reverseStr(testString)),
			},
			want: [2]BigInt{
				{
					Digits: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					IsNeg:  false,
				},
				{
					Digits: []int{49374, 47806, 61453, 51966, 48879, 57005, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					IsNeg:  false,
				},
			},
		},
		{
			name: "SmallNeg/BigPos",
			args: args{
				x: BiFromHex("-" + testString[0:24]),
				y: BiFromHex(reverseStr(testString)),
			},
			want: [2]BigInt{
				{
					Digits: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					IsNeg:  true,
				},
				{
					Digits: []int{6671, 17453, 65439, 1296, 11452, 3679, 60591, 56043, 0, 0, 0, 0, 0, 0, 0, 0},
					IsNeg:  false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiDivideModulo(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biDivideModulo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biDivide(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiDivide(testBigInt, testBigInt)
	}
}

func Test_biDivide(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex(reverseStr(testString)),
			},
			want: bigOne,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiDivide(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biModulo(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiModulo(testBigInt, testBigInt)
	}
}

func Test_biModulo(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex(reverseStr(testString)),
			},
			want: BigInt{
				Digits: []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961},
				IsNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiModulo(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biModulo() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biMultiplyMod(t *testing.T) {}

// func Test_biMultiplyMod(t *testing.T) {}

// func Test_biPow(t *testing.T) {}

// func Test_biPowMod(t *testing.T) {}
