package gui

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw() {
	fontSize := int32(20)

	rg.SetStyleColor(rg.ButtonDefaultTextColor, rl.Black)
	rg.SetStyleColor(rg.ButtonHoverTextColor, rl.Black)
	rg.SetStyleProperty(rg.GlobalTextFontsize, int64(fontSize))

	DrawConfigButtons()
	DrawShapeButtons()
}
