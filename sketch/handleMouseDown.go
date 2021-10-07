package sketch

import (
	"github.com/PlebusSupremus1234/Game-of-Life/global"
	"math"
)

func HandleMouseDown(x int32, y int32) {
	diffX := math.Abs(float64(x - prevMouseX))
	diffY := math.Abs(float64(y - prevMouseY))

	if diffX < 10 && diffY < 10 {
		col := int(math.Floor(float64(y / 10)))
		row := int(math.Floor(float64(x / 10)))

		if global.Tool == 0 {
			global.Grid[col][row] = 1
		} else {
			global.Grid[col][row] = 0
		}
	} else {
		// Find y = mx + b
		var m float64 // Slope
		if x-prevMouseX == 0 {
			m = 100000 // Close enough to an infinite value
		} else {
			m = float64(y-prevMouseY) / float64(x-prevMouseX)
		}

		b := float64(y) - float64(x)*m // Y intercept

		// Find side values
		minX := math.Min(float64(prevMouseX), float64(x))
		maxX := math.Max(float64(prevMouseX), float64(x))

		minY := math.Min(float64(prevMouseY), float64(y))
		maxY := math.Max(float64(prevMouseY), float64(y))

		// Loop through points
		for i := minY; i <= maxY; i += 10 {
			for j := minX; j <= maxX; j += 10 {
				// Middle of the cell
				xVal := j + 5
				yVal := i + 5

				// Calculate the distance from the line
				var dist float64
				if m == 0 {
					dist = math.Abs(-yVal + b)
				} else {
					dist = math.Abs(m*xVal-yVal+b) / math.Sqrt(m*m+1)
				}

				// Simplified expression for sqrt(10^2 + 10^2) / 2
				if dist < 5*math.Sqrt(2) {
					col := int(math.Floor(yVal / 10))
					row := int(math.Floor(xVal / 10))

					if col >= 0 && col < 100 && row >= 0 && row < 100 {
						if global.Tool == 0 {
							global.Grid[col][row] = 1
						} else {
							global.Grid[col][row] = 0
						}
					}
				}
			}
		}
	}

	prevMouseX = x
	prevMouseY = y
}
