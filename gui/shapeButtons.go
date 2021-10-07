package gui

import (
	"github.com/PlebusSupremus1234/Game-of-Life/global"
	"github.com/PlebusSupremus1234/Game-of-Life/shapes"
)

func DrawShapeButtons() {
	padding := int32(50)
	fontSize := int32(20)

	if DrawButton("Hammerhead Spaceship", fontSize, padding, 600, 60) {
		global.IsPlacing = !global.IsPlacing
		global.SelectedShape = shapes.Hammerhead
	}

	if DrawButton("Gosper Glider Gun", fontSize, padding, 670, 60) {
		global.IsPlacing = !global.IsPlacing
		global.SelectedShape = shapes.GliderGun
	}

	if DrawButton("Glider Spaceship", fontSize, padding, 740, 60) {
		global.IsPlacing = !global.IsPlacing
		global.SelectedShape = shapes.Glider
	}

	if DrawButton("R Pentomino", fontSize, padding, 810, 60) {
		global.IsPlacing = !global.IsPlacing
		global.SelectedShape = shapes.RPentomino
	}
}
