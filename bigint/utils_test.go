package bigint

import (
	"reflect"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		{
			args: args{
				s: "deadbeef",
			},
			want: "feebdaed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digitToHex(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitToHex(tt.args.n); got != tt.want {
				t.Errorf("digitToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_charToHex(t *testing.T) {
	type args struct {
		c rune
	}
	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			args: args{
				c: rune("0"[0]),
			},
			want: 0,
		},
		{
			args: args{
				c: rune("9"[0]),
			},
			want: 9,
		},
		{
			args: args{
				c: rune("f"[0]),
			},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := charToHex(tt.args.c); got != tt.want {
				t.Errorf("charToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexToDigit(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			args: args{
				s: "dead",
			},
			want: 57005,
		},
		{
			args: args{
				s: "beef",
			},
			want: 48879,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexToDigit(tt.args.s); got != tt.want {
				t.Errorf("hexToDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digitToBytes(t *testing.T) {
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitToBytes(tt.args.n); got != tt.want {
				t.Errorf("digitToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrayCopy(t *testing.T) {
	type args struct {
		src       []int
		srcStart  int
		dest      []int
		destStart int
		n         int
	}
	type test struct {
		name string
		args args
		want []int
	}

	tests := []test{
		{
			args: args{
				src:       []int{13, 14, 10, 13, 11, 14, 14, 15},
				srcStart:  2,
				dest:      []int{0, 0, 0, 0, 0, 0, 0, 0},
				destStart: 2,
				n:         4,
			},
			want: []int{0, 0, 10, 13, 11, 14, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrayCopy(tt.args.src, tt.args.srcStart, tt.args.dest, tt.args.destStart, tt.args.n)
			if !reflect.DeepEqual(tt.args.dest, tt.want) {
				t.Errorf("arrayCopy() = %v, want %v", tt.args.dest, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		x int
		y int
	}
	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			name: "x",
			args: args{
				x: 2,
				y: 1,
			},
			want: 2,
		},
		{
			name: "y",
			args: args{
				x: 1,
				y: 2,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		x int
		y int
	}
	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			name: "x",
			args: args{
				x: 1,
				y: 2,
			},
			want: 1,
		},
		{
			name: "y",
			args: args{
				x: 2,
				y: 1,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ceil(t *testing.T) {
	type args struct {
		x int
		y int
	}
	type test struct {
		name string
		args args
		want int
	}

	tests := []test{
		{
			name: "x",
			args: args{
				x: 1,
				y: 2,
			},
			want: 1,
		},
		{
			name: "y",
			args: args{
				x: 2,
				y: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ceil(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}
