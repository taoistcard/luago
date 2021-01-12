package number

import "math"

//IFloorDiv 整数整除
func IFloorDiv(a, b int64) int64 {
	if ((a > 0) && (b > 0)) || ((a < 0) && (b < 0)) || (a%b == 0) {
		return a / b
	}
	return a/b - 1
}

//FFloorDiv 浮点数整除
func FFloorDiv(a, b float64) float64 {
	return math.Floor(a / b)
}

func IMod(a, b int64) int64 {
	return a - IFloorDiv(a, b)*b
}

func FMod(a, b float64) float64 {
	return a - math.Floor(a/b)*b
}

func ShiftLeft(a, n int64) int64 {
	if n >= 0 {
		return a << uint64(n)
	} else {
		return ShiftRight(a, -n)
	}
}

//有符号整数右移，高位会填充1，所以这里先转出无符号整数再移位
func ShiftRight(a, n int64) int64 {
	if n >= 0 {
		return int64(uint64(a) >> uint64(n))
	} else {
		return ShiftLeft(a, -n)
	}
}

func FloatToInteger(f float64) (int64, bool) {
	i := int64(f)
	return i, float64(i) == f
}
