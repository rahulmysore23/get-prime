package utilities

import "math"

func Sqrt(n int64) int64 {
	return int64(math.Floor(float64(n) / 2))
}
