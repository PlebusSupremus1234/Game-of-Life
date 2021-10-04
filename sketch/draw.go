package sketch

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func Draw(array [100][100]int, paused bool) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	w := 10
	guideLines := false

	for i, subarray := range array {
		for j, element := range subarray {
			if element == 1 {
				rl.DrawRectangle(int32(j*w), int32(i*w), int32(w), int32(w), rl.White)
			}

			if guideLines {
				rl.DrawLine(int32(j*10), 0, int32(j*10), 1000, rl.Gray) // Vertical
			}
		}

		if guideLines {
			rl.DrawLine(0, int32(i*10), 1000, int32(i*10), rl.Gray) // Horizontal
		}
	}

	rl.DrawRectangle(1000, 0, 300, 1000, rl.Gray)

	CenterText("Game of Life", 42, 1000, 15, 300)
	CenterText("By PlebusSupremus1234", 20, 1000, 65, 300)

	text := "Paused"
	if !paused {
		text = "Playing"
	}
	LeftText("Status: "+text, 30, 1010, 110)

	text = "Pencil"
	if tool == 1 {
		text = "Eraser"
	}
	LeftText("Tool: "+text, 30, 1010, 150)

	text = "Off"
	if guideLines {
		text = "On"
	}
	LeftText("Guidelines: "+text, 30, 1010, 190)

	rl.EndDrawing()
}
