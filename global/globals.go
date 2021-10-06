package global

import "PlebusSupremus1234/game_of_life/shapes"

var Width = int32(1000)
var ConfigWidth = int32(300)
var CellWidth = int32(10)

var Grid = [100][100]int{}
var Paused = true
var GuideLines = false
var Tool = 0

// 0 = draw
// 1 = erase

var IsPlacing = false
var SelectedShape []shapes.Coord
