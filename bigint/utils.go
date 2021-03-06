package bigint

func reverseStr(s string) string {
	result := ""
	for i := len(s) - 1; i > -1; i-- {
		result += string(s[i])
	}
	return result
}

var hexToChar = []string{
	"0", "1", "2", "3",
	"4", "5", "6", "7",
	"8", "9", "a", "b",
	"c", "d", "e", "f",
}

var hexatrigesimalToChar = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
}

func digitToHex(n int) string {
	mask := 0xF
	var result string
	for i := 0; i < 4; i++ {
		result += hexToChar[n&mask]
		n >>= 4
	}
	return reverseStr(result)
}

func charToHex(c rune) int {
	ZERO := 48
	NINE := ZERO + 9
	littleA := 97
	littleZ := littleA + 25
	bigA := 65
	bigZ := 65 + 25
	result := 0

	i := int(c)
	if i >= ZERO && i <= NINE {
		result = i - ZERO
	} else if i >= bigA && i <= bigZ {
		result = 10 + i - bigA
	} else if i >= littleA && i <= littleZ {
		result = 10 + i - littleA
	} else {
		result = 0
	}

	return result
}

func hexToDigit(s string) int {
	result := 0
	sl := min(len(s), 4)
	for i := 0; i < sl; i++ {
		result = result << 4
		result = result | charToHex(rune(s[i]))
	}
	return result
}

func digitToBytes(n int) string {
	// NOTE: This returns invalid UTF-8 character values.
	c1 := string(rune(n & 0xFF))
	n >>= 8
	c2 := string(rune(n & 0xFF))
	return c2 + c1
}

func arrayCopy(src []int, srcStart int, dest []int, destStart int, n int) {
	m := min(srcStart+n, len(src))
	for i, j := srcStart, destStart; i < m; i, j = i+1, j+1 {
		dest[j] = src[i]
	}
}

var highBitMasks = []int{
	0x0000, 0x8000, 0xC000, 0xE000,
	0xF000, 0xF800, 0xFC00, 0xFE00,
	0xFF00, 0xFF80, 0xFFC0, 0xFFE0,
	0xFFF0, 0xFFF8, 0xFFFC, 0xFFFE,
	0xFFFF,
}

var lowBitMasks = []int{
	0x0000, 0x0001, 0x0003, 0x0007,
	0x000F, 0x001F, 0x003F, 0x007F,
	0x00FF, 0x01FF, 0x03FF, 0x07FF,
	0x0FFF, 0x1FFF, 0x3FFF, 0x7FFF,
	0xFFFF,
}

// max returns the larger of x or y.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// min returns the smaller of x or y.
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func ceil(x, y int) int {
	result := (x + y - 1) / y
	return result
}
