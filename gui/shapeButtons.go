package gui

import (
	"PlebusSupremus1234/game_of_life/global"
	"PlebusSupremus1234/game_of_life/shapes"
)

func DrawShapeButtons() {
	padding := int32(50)
	fontSize := int32(20)

	text := "R Pentomino"

	if DrawButton(text, fontSize, padding, 600, 60) {
		global.IsPlacing = !global.IsPlacing
		global.SelectedShape = shapes.RPentomino
	}
}
