package sketch

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

var array = [100][100]int{}
var paused = true
var tool = 0

// 0 = draw
// 1 = erase

var prevMouseX = int32(-1)
var prevMouseY = int32(-1)

func Run() {
	width := 1000

	if rl.IsKeyPressed(rl.KeySpace) {
		paused = !paused
	}

	if rl.IsKeyPressed(rl.KeyD) {
		tool = 0
	}
	if rl.IsKeyPressed(rl.KeyE) {
		tool = 1
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		prevMouseX = rl.GetMouseX()
		prevMouseY = rl.GetMouseY()
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		x := rl.GetMouseX()
		y := rl.GetMouseY()

		diffX := math.Abs(float64(x - prevMouseX))
		diffY := math.Abs(float64(y - prevMouseY))

		if x >= 0 && x < int32(width) && y >= 0 && y < int32(width) {
			if diffX < 10 && diffY < 10 {
				col := int(math.Floor(float64(y / 10)))
				row := int(math.Floor(float64(x / 10)))

				if tool == 0 {
					array[col][row] = 1
				} else {
					array[col][row] = 0
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
								if tool == 0 {
									array[col][row] = 1
								} else {
									array[col][row] = 0
								}
							}
						}
					}
				}
			}

			prevMouseX = x
			prevMouseY = y
		}
	}

	if !paused {
		Update(&array)
	}

	Draw(array, paused)
}
