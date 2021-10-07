package gui

import (
	"github.com/PlebusSupremus1234/Game-of-Life/global"
	rg "github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

func DrawButton(text string, fontSize int32, padding int32, y float32, h float32) bool {
	textWidth := rl.MeasureText(text, fontSize)

	rect := rl.Rectangle{
		X:      float32(global.Width + (global.ConfigWidth-textWidth-padding)/2),
		Y:      y,
		Width:  float32(textWidth + padding),
		Height: h,
	}

	return rg.Button(rect, text)
}
