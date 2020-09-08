package bigint

import (
	"reflect"
	"testing"
)

// Used to store benchmark results
var rBigInt BigInt
var rString string
var rInt int

var testRunes = []rune{222, 173, 190, 239, 202, 254, 240, 13, 186, 190, 192, 222, 250, 206, 190, 173}
var testString = "dead" + "beef" + "cafe" + "f00d" + "babe" + "c0de" + "face" + "bead"
var testBigInt = BigInt{
	Digits: []int{48813, 64206, 49374, 47806, 61453, 51966, 48879, 57005}, // bead, face, c0de, babe, f00d, cafe, beef, dead
	IsNeg:  false,
}

func Test_SetMaxDigits(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMaxDigits(tt.args.value)
		})
	}
}

func Test_NewBigInt(t *testing.T) {
	type args struct {
		flag bool
	}
	type test struct {
		name string
		args args
		want BigInt
	}

	tests := []test{
		{
			args: args{
				flag: false,
			},
			want: BigInt{
				Digits: []int{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
				IsNeg: false,
			},
		},
		{
			args: args{
				flag: true,
			},
			want: BigInt{
				Digits: []int{},
				IsNeg:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_BiFromDecimal(t *testing.T) {}

func Benchmark_BiCopy(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiCopy(testBigInt)
	}
}

func Test_BiCopy(t *testing.T) {
	type args struct {
		bi BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

// func Test_BiFromNumber(t *testing.T) {}

// func Test_BiToString(t *testing.T) {}

// func Test_BiToDecimal(t *testing.T) {}

func Benchmark_BiToHex(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rString = BiToHex(testBigInt)
	}
}

func Test_BiToHex(t *testing.T) {
	type args struct {
		x BigInt
	}
	type test struct {
		name string
		args args
		want string
	}

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

func Benchmark_BiFromHex(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiFromHex(testString)
	}
}

func Test_BiFromHex(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

// func Test_BiFromString(t *testing.T) {}

func Benchmark_BiToBytes(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rString = BiToBytes(testBigInt)
	}
}

func Test_BiToBytes(t *testing.T) {
	type args struct {
		x BigInt
	}
	type test struct {
		name string
		args args
		want string
	}

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

// func Test_BiDump(t *testing.T) {}

func Benchmark_BiAdd(b *testing.B) {
	x := BiFromHex("deadbeef")
	y := BiFromHex("beefdead")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiAdd(x, y)
	}
}

func Test_BiAdd(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiSubtract(b *testing.B) {
	x := BiFromHex(testString)
	y := BiFromHex(reverseStr(testString))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiSubtract(x, y)
	}
}

func Test_BiSubtract(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiHighIndex(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = BiHighIndex(testBigInt)
	}
}

func Test_BiHighIndex(t *testing.T) {
	type args struct {
		x BigInt
	}
	type test struct {
		name string
		args args
		want int
	}

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

func Benchmark_BiNumBits(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = BiNumBits(testBigInt)
	}
}

func Test_BiNumBits(t *testing.T) {
	type args struct {
		x BigInt
	}
	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			args: args{
				x: testBigInt,
			},
			want: 128,
		},
		{
			args: args{
				x: BiFromHex(testString[0:31]),
			},
			want: 124,
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

func Benchmark_BiMultiply(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiMultiply(testBigInt, testBigInt)
	}
}

func Test_BiMultiply(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiMultiplyDigit(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiMultiplyDigit(testBigInt, 16)
	}
}

func Test_BiMultiplyDigit(t *testing.T) {
	type args struct {
		x BigInt
		y int
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiShiftLeft(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiShiftLeft(testBigInt, 16)
	}
}

func Test_BiShiftLeft(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiShiftRight(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiShiftRight(testBigInt, 16)
	}
}

func Test_BiShiftRight(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiMultiplyByRadixPower(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiMultiplyByRadixPower(testBigInt, 3)
	}
}

func Test_BiMultiplyByRadixPower(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiDivideByRadixPower(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiDivideByRadixPower(testBigInt, 3)
	}
}

func Test_BiDivideByRadixPower(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiModuloByRadixPower(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiModuloByRadixPower(testBigInt, 3)
	}
}

func Test_BiModuloByRadixPower(t *testing.T) {
	type args struct {
		x BigInt
		n int
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiCompare(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rInt = BiCompare(testBigInt, testBigInt)
	}
}

func Test_BiCompare(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want int
	}

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

func Benchmark_BiDivideModulo(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiDivideModulo(testBigInt, testBigInt)[0]
	}
}

func Test_BiDivideModulo(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want [2]BigInt
	}

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
			name: "Big/Small",
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex(reverseStr(testString[0:24])),
			},
			want: [2]BigInt{
				{
					Digits: []int{46954, 61562, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					IsNeg:  false,
				},
				{
					Digits: []int{44427, 20148, 16222, 51298, 28899, 1085, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
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
					Digits: bigOne.Digits,
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

func Benchmark_BiDivide(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiDivide(testBigInt, testBigInt)
	}
}

func Test_BiDivide(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}

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

func Benchmark_BiModulo(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		rBigInt = BiModulo(testBigInt, testBigInt)
	}
}

func Test_BiModulo(t *testing.T) {
	type args struct {
		x BigInt
		y BigInt
	}
	type test struct {
		name string
		args args
		want BigInt
	}

	value := []int{58304, 64482, 53553, 60078, 1121, 56818, 53823, 961}
	digits := bigZero.Digits
	for i, v := range value {
		digits[i] = v
	}

	tests := []test{
		{
			args: args{
				x: BiFromHex(testString),
				y: BiFromHex(reverseStr(testString)),
			},
			want: BigInt{
				Digits: digits,
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

// func Test_BiMultiplyMod(t *testing.T) {}

// func Test_BiMultiplyMod(t *testing.T) {}

// func Test_BiPow(t *testing.T) {}

// func Test_BiPowMod(t *testing.T) {}
