package main

import (
	"PlebusSupremus1234/game_of_life/sketch"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	width := 1000

	rl.InitWindow(int32(width)+300, int32(width), "Conway's Game of Life by PlebusSupremus1234")

	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		sketch.Run()
	}

	rl.CloseWindow()
}
