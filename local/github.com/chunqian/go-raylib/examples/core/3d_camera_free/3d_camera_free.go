package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d camera free")
	defer rl.CloseWindow()

	camera := rl.NewCamera3D(
		rl.NewVector3(10, 10, 10),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	cubePosition := rl.NewVector3(0, 0, 0)

	rl.SetCameraMode(rl.Camera(camera), int32(rl.CAMERA_FREE))
	rl.SetCameraSmoothZoomControl(int32(rl.KEY_LEFT_SHIFT))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera((*rl.Camera)(&camera))

		if rl.IsKeyDown(int32(rl.KEY_Z)) {
			camera.Target = rl.NewVector3(0, 0, 0)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 320, 133, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 320, 133, rl.Blue)

		rl.DrawText("Free camera default controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Mouse Wheel to Zoom in-out", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Mouse Wheel Pressed to Pan", 40, 60, 10, rl.DarkGray)
		rl.DrawText("- Alt + Mouse Wheel Pressed to Rotate", 40, 80, 10, rl.DarkGray)
		rl.DrawText("- Alt + Left Shift + Mouse Wheel Pressed for Smooth Zoom", 40, 100, 10, rl.DarkGray)
		rl.DrawText("- Z to zoom to (0, 0, 0)", 40, 120, 10, rl.DarkGray)

		rl.EndDrawing()
	}
}
