package gui

import (
	"github.com/PlebusSupremus1234/Game-of-Life/global"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw() {
	rect := rl.Rectangle{
		X:      float32(global.Width + 20),
		Y:      270,
		Width:  float32(global.ConfigWidth - 40),
		Height: 20,
	}
	global.TargetFPS = int32(rg.Slider(rect, float32(global.TargetFPS), 1, 40))

	fontSize := int32(20)

	rg.SetStyleColor(rg.ButtonDefaultTextColor, rl.Black)
	rg.SetStyleColor(rg.ButtonHoverTextColor, rl.Black)
	rg.SetStyleProperty(rg.GlobalTextFontsize, int64(fontSize))

	DrawConfigButtons()
	DrawShapeButtons()
}
