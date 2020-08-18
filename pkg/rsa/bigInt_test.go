package rsa

import (
	"reflect"
	"testing"
)

// Used to store benchmark results
var rBigInt bigInt
var rString string
var rInt int

var testRunes = []rune{222, 173, 190, 239, 202, 254, 240, 13, 186, 190, 192, 222, 250, 206, 190, 173}
var testString = "deadbeefcafef00dbabec0defacebead"
var testBigInt = bigInt{
	digits: []int{48813, 64206, 49374, 47806, 61453, 51966, 48879, 57005}, // bead, face, c0de, babe, f00d, cafe, beef, dead
	isNeg:  false,
}

// func Test_biFromDecimal(t *testing.T) {}

func Benchmark_biCopy(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biCopy(testBigInt)
	}
}

func Test_biCopy(t *testing.T) {
	type args struct {
		bi bigInt
	}
	type test struct {
		name string
		args args
		want bigInt
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
			if got := biCopy(tt.args.bi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biFromNumber(t *testing.T) {}

// func Benchmark_biToString(b *testing.B) {}

func Test_biToString(t *testing.T) {
	type args struct {
		x     bigInt
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
			if got := biToString(tt.args.x, tt.args.radix); got != tt.want {
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
		rString = biToHex(testBigInt)
	}
}

func Test_biToHex(t *testing.T) {
	type args struct {
		x bigInt
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
			if got := biToHex(tt.args.x); got != tt.want {
				t.Errorf("biToHex() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Benchmark_biFromHex(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biFromHex(testString)
	}
}

func Test_biFromHex(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want bigInt
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
			if got := biFromHex(tt.args.s); !reflect.DeepEqual(got, tt.want) {
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
		rString = biToBytes(testBigInt)
	}
}

func Test_biToBytes(t *testing.T) {
	type args struct {
		x bigInt
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
			if got := biToBytes(tt.args.x); got != tt.want {
				t.Errorf("biToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biDump(t *testing.T) {}

func Benchmark_biAdd(b *testing.B) {
	SetMaxDigits(16)
	x := biFromHex("deadbeef")
	y := biFromHex("beefdead")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biAdd(x, y)
	}
}

func Test_biAdd(t *testing.T) {
	type args struct {
		x bigInt
		y bigInt
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: biFromHex("deadbeef"),
				y: biFromHex("beefdead"),
			},
			want: bigInt{
				digits: []int{40348, 40349, 1, 0, 0, 0, 0, 0},
				isNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biAdd(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biSubtract(b *testing.B) {
	SetMaxDigits(16)
	x := biFromHex(testString)
	y := biFromHex(reverseStr(testString))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biSubtract(x, y)
	}
}

func Test_biSubtract(t *testing.T) {
	type args struct {
		x bigInt
		y bigInt
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: biFromHex(testString),
				y: biFromHex(reverseStr(testString)),
			},
			want: bigInt{
				digits: []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961},
				isNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biSubtract(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biSubtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biHighIndex(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = biHighIndex(testBigInt)
	}
}

func Test_biHighIndex(t *testing.T) {
	type args struct {
		x bigInt
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
			if got := biHighIndex(tt.args.x); got != tt.want {
				t.Errorf("biHighIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biNumBits(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = biNumBits(testBigInt)
	}
}

func Test_biNumBits(t *testing.T) {
	type args struct {
		x bigInt
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
			if got := biNumBits(tt.args.x); got != tt.want {
				t.Errorf("biNumBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biMultiply(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biMultiply(testBigInt, testBigInt)
	}
}

func Test_biMultiply(t *testing.T) {
	type args struct {
		x bigInt
		y bigInt
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(16)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				y: testBigInt,
			},
			want: bigInt{
				digits: []int{
					16617, 20593, 42861, 32538, 45046, 25231, 13697, 47718,
					1165, 9698, 10208, 22465, 22231, 33427, 52499, 49585,
				},
				isNeg: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biMultiply(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biMultiplyDigit(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biMultiplyDigit(testBigInt, 16)
	}
}

func Test_biMultiplyDigit(t *testing.T) {
	type args struct {
		x bigInt
		y int
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(16)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				y: 16,
			},
			want: bigInt{
				digits: []int{
					60112, 44267, 3567, 44012, 219, 45039, 61180, 60123,
					13, 0, 0, 0, 0, 0, 0, 0,
				},
				isNeg: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biMultiplyDigit(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biMultiplyDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biShiftLeft(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biShiftLeft(testBigInt, 16)
	}
}

func Test_biShiftLeft(t *testing.T) {
	type args struct {
		x bigInt
		n int
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 16,
			},
			want: bigInt{
				digits: []int{0, 48813, 64206, 49374, 47806, 61453, 51966, 48879},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biShiftLeft(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biShiftLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biShiftRight(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biShiftRight(testBigInt, 16)
	}
}

func Test_biShiftRight(t *testing.T) {
	type args struct {
		x bigInt
		n int
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 16,
			},
			want: bigInt{
				digits: []int{64206, 49374, 47806, 61453, 51966, 48879, 57005, 0},
				isNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biShiftRight(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biShiftRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biMultiplyByRadixPower(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biMultiplyByRadixPower(testBigInt, 3)
	}
}

func Test_biMultiplyByRadixPower(t *testing.T) {
	type args struct {
		x bigInt
		n int
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 3,
			},
			want: bigInt{
				digits: []int{0, 0, 0, 48813, 64206, 49374, 47806, 61453},
				isNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biMultiplyByRadixPower(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biMultiplyByRadixPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biDivideByRadixPower(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biDivideByRadixPower(testBigInt, 3)
	}
}

func Test_biDivideByRadixPower(t *testing.T) {
	type args struct {
		x bigInt
		n int
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 3,
			},
			want: bigInt{
				digits: []int{47806, 61453, 51966, 48879, 57005, 0, 0, 0},
				isNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biDivideByRadixPower(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biDivideByRadixPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biModuloByRadixPower(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biModuloByRadixPower(testBigInt, 3)
	}
}

func Test_biModuloByRadixPower(t *testing.T) {
	type args struct {
		x bigInt
		n int
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: testBigInt,
				n: 3,
			},
			want: bigInt{
				digits: []int{48813, 64206, 49374, 0, 0, 0, 0, 0},
				isNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biModuloByRadixPower(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biModuloByRadixPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biCompare(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = biCompare(testBigInt, testBigInt)
	}
}

func Test_biCompare(t *testing.T) {
	type args struct {
		x bigInt
		y bigInt
	}
	type test struct {
		name string
		args args
		want int
	}
	SetMaxDigits(8)

	w, x, y, z := "-bead", "-face", "dead", "beef"
	biw, bix, biy, biz := biFromHex(w), biFromHex(x), biFromHex(y), biFromHex(z)

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
			if got := biCompare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("biCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biDivideModulo(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biDivideModulo(testBigInt, testBigInt)[0]
	}
}

func Test_biDivideModulo(t *testing.T) {
	type args struct {
		x bigInt
		y bigInt
	}
	type test struct {
		name string
		args args
		want [2]bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: biFromHex(testString),
				y: biFromHex(reverseStr(testString)),
			},
			want: [2]bigInt{
				bigOne,
				{
					digits: []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961},
					isNeg:  false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biDivideModulo(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biDivideModulo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biDivide(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biDivide(testBigInt, testBigInt)
	}
}

func Test_biDivide(t *testing.T) {
	type args struct {
		x bigInt
		y bigInt
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: biFromHex(testString),
				y: biFromHex(reverseStr(testString)),
			},
			want: bigOne,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biDivide(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_biModulo(b *testing.B) {
	SetMaxDigits(16)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = biModulo(testBigInt, testBigInt)
	}
}

func Test_biModulo(t *testing.T) {
	type args struct {
		x bigInt
		y bigInt
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		{
			args: args{
				x: biFromHex(testString),
				y: biFromHex(reverseStr(testString)),
			},
			want: bigInt{
				digits: []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961},
				isNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biModulo(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biModulo() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biMultiplyMod(t *testing.T) {}

func Test_biMultiplyMod(t *testing.T) {
	type args struct {
		x bigInt
		y bigInt
		m bigInt
	}
	type test struct {
		name string
		args args
		want bigInt
	}
	SetMaxDigits(8)

	tests := []test{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := biMultiplyMod(tt.args.x, tt.args.y, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biMultiplyMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_biPow(t *testing.T) {}

// func Test_biPowMod(t *testing.T) {}
