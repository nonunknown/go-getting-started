package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"fmt"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - input mouse wheel")
	defer rl.CloseWindow()

	boxPositionY := screenHeight/2 - 40
	scrollSpeed := int32(4)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		boxPositionY -= int32(rl.GetMouseWheelMove()) * scrollSpeed

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangle(screenWidth/2-40, boxPositionY, 80, 80, rl.Maroon)
		rl.DrawText(fmt.Sprintf("Box position Y: %d", boxPositionY), 10, 40, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
