package utilities

import "math"

// For test - Remove after adding UI
const DefaultHtml = `<html>
<body>
<div style="text-align: center; margin-top: 100px;">
<h1>Get Prime</h1>
</div>
</body>
</html>`

func Sqrt(n int64) int64 {
	return int64(math.Floor(float64(n) / 2))
}
