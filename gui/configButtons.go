package gui

import "PlebusSupremus1234/game_of_life/global"

func DrawConfigButtons() {
	padding := int32(50)
	fontSize := int32(20)

	text := "Pause (Space)"
	if global.Paused {
		text = "Play (Space)"
	}

	if DrawButton(text, fontSize, padding, 250, 60) {
		global.Paused = !global.Paused
	}

	text = "Switch to Eraser (E)"
	if global.Tool == 1 {
		text = "Switch to Pencil (P)"
	}

	if DrawButton(text, fontSize, padding, 320, 60) {
		if global.Tool == 0 {
			global.Tool = 1
		} else {
			global.Tool = 0
		}
	}

	text = "Toggle Guidelines"

	if DrawButton(text, fontSize, padding, 390, 60) {
		global.GuideLines = !global.GuideLines
	}

	text = "Clear Grid"

	if DrawButton(text, fontSize, padding, 460, 60) {
		global.Grid = [100][100]int{}
	}
}
