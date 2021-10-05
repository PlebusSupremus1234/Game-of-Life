package gui

import (
	"PlebusSupremus1234/game_of_life/global"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw() {
	fontSize := int32(20)
	padding := int32(50)

	rg.SetStyleColor(rg.ButtonDefaultTextColor, rl.Black)
	rg.SetStyleColor(rg.ButtonHoverTextColor, rl.Black)
	rg.SetStyleProperty(rg.GlobalTextFontsize, int64(fontSize))

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
}
