package sketch

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func Draw(array [100][100]int) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	w := 10

	for i, subarray := range array {
		for j, element := range subarray {
			color := rl.Black
			if element == 1 { color = rl.White }

			rl.DrawRectangle(int32(j * w), int32(i * w), int32(w), int32(w), color)
		}
	}

	rl.EndDrawing()
}