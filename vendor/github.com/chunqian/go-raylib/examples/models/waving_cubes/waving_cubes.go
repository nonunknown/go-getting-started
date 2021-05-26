package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"math"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - waving cubes")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(30.0, 20.0, 30.0),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		70.0,
		int32(rl.CAMERA_PERSPECTIVE),
	)

	numBlocks := 15

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		time := rl.GetTime()

		scale := (2.0 + math.Sin(time)) * 0.7

		cameraTime := time * 0.3
		camera.Position.X = float32(math.Cos(cameraTime) * 40.0)
		camera.Position.Z = float32(math.Sin(cameraTime) * 40.0)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawGrid(10, 5.0)

		for x := 0; x < numBlocks; x++ {

			for y := 0; y < numBlocks; y++ {

				for z := 0; z < numBlocks; z++ {

					blockScale := float64(x+y+z) / 30.0

					scatter := math.Sin(blockScale*20.0 + time*4.0)

					cubePos := rl.NewVector3(
						float32((float64(x)-float64(numBlocks)/2)*(scale*3.0)+scatter),
						float32((float64(y)-float64(numBlocks)/2)*(scale*3.0)+scatter),
						float32((float64(z)-float64(numBlocks)/2)*(scale*3.0)+scatter),
					)

					cubeColor := rl.ColorFromHSV(float32(((x+y+z)*18)%360), 0.75, 0.9)
					cubeSize := float32((2.4 - scale) * blockScale)

					rl.DrawCube(cubePos, cubeSize, cubeSize, cubeSize, cubeColor)
				}
			}
		}

		rl.EndMode3D()

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
