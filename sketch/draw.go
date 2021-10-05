package sketch

import (
	"PlebusSupremus1234/game_of_life/global"
	"PlebusSupremus1234/game_of_life/gui"
	"github.com/gen2brain/raylib-go/raylib"
)

func Draw() {

	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	w := global.CellWidth

	for i, subarray := range global.Grid {
		for j, element := range subarray {
			if element == 1 {
				rl.DrawRectangle(int32(j*w), int32(i*w), int32(w), int32(w), rl.White)
			}

			if global.GuideLines {
				rl.DrawLine(int32(j*10), 0, int32(j*10), 1000, rl.Gray) // Vertical
			}
		}

		if global.GuideLines {
			rl.DrawLine(0, int32(i*10), 1000, int32(i*10), rl.Gray) // Horizontal
		}
	}

	rl.DrawRectangle(1000, 0, 300, 1000, rl.Gray)

	CenterText("Game of Life", 42, 1000, 15, 300)
	CenterText("By PlebusSupremus1234", 20, 1000, 65, 300)

	text := "Playing"
	if global.Paused {
		text = "Paused"
	}
	LeftText("Status: "+text, 30, 1010, 110)

	text = "Pencil"
	if global.Tool == 1 {
		text = "Eraser"
	}
	LeftText("Tool: "+text, 30, 1010, 150)

	text = "Off"
	if global.GuideLines {
		text = "On"
	}
	LeftText("Guidelines: "+text, 30, 1010, 190)

	// Gui
	gui.Draw()

	rl.EndDrawing()
}
