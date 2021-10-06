package sketch

import (
	"PlebusSupremus1234/game_of_life/global"
	"github.com/gen2brain/raylib-go/raylib"
)

var prevMouseX = int32(-1)
var prevMouseY = int32(-1)

func Run() {
	width := 1000

	if rl.IsKeyPressed(rl.KeySpace) {
		global.Paused = !global.Paused
	}

	if rl.IsKeyPressed(rl.KeyP) {
		global.Tool = 0
	}
	if rl.IsKeyPressed(rl.KeyE) {
		global.Tool = 1
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		prevMouseX = rl.GetMouseX()
		prevMouseY = rl.GetMouseY()

		if global.IsPlacing && prevMouseX >= 0 && prevMouseX < 1000 && prevMouseY >= 0 && prevMouseY < 1000 {
			global.IsPlacing = false

			x := prevMouseX / 10
			y := prevMouseY / 10

			for _, element := range global.SelectedShape {
				col := y + int32(element.Y)
				row := x + int32(element.X)

				if col >= 0 && col < 100 && row >= 0 && row < 100 {
					global.Grid[col][row] = 1
				}
			}
		}
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		x := rl.GetMouseX()
		y := rl.GetMouseY()

		if x >= 0 && x < int32(width) && y >= 0 && y < int32(width) {
			HandleMouseDown(x, y)
		}
	}

	if !global.Paused {
		Update()
	}

	Draw()
}
