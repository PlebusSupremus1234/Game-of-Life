package main

import (
	"PlebusSupremus1234/game_of_life/sketch"
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

func main() {
	width := 1000

	rl.InitWindow(int32(width), int32(width), "Conway's Game of Life by PlebusSupremus1234")

	rl.SetTargetFPS(60)

	array := [100][100]int {}

	paused := true

	tool := 0
	// 0 = draw
	// 1 = erase

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) { paused = !paused }

		if rl.IsKeyPressed(rl.KeyD) { tool = 0 }
		if rl.IsKeyPressed(rl.KeyE) { tool = 1 }

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			x := rl.GetMouseX()
			y := rl.GetMouseY()

			if x >= 0 && x <= int32(width) && y >= 0 && y <= int32(width) {
				col := int(math.Floor(float64(y / 10)))
				row := int(math.Floor(float64(x / 10)))

				if tool == 0 {
					array[col][row] = 1
				} else {
					array[col][row] = 0
				}
			}
		}

		if !paused { sketch.Update(&array) }
		sketch.Draw(array)
	}

	rl.CloseWindow()
}
