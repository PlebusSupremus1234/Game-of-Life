package sketch

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func LeftText(text string, fontSize int, left int32, top int32) {
	rl.DrawText(text, left, top, int32(fontSize), rl.Black)
}

func CenterText(text string, fontSize int, left int32, top int32, width int32) {
	length := rl.MeasureText(text, int32(fontSize))
	rl.DrawText(text, left+(width-length)/2, top, int32(fontSize), rl.Black)
}
