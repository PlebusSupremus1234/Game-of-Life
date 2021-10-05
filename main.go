package main

import (
	"PlebusSupremus1234/game_of_life/global"
	"PlebusSupremus1234/game_of_life/sketch"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(global.Width+global.ConfigWidth, global.Width, "Conway's Game of Life by PlebusSupremus1234")

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		sketch.Run()
	}

	rl.CloseWindow()
}
