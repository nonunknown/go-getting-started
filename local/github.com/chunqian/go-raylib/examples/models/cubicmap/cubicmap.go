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

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - cubesmap loading and drawing")
	defer rl.CloseWindow()

	camera := rl.NewCamera(
		rl.NewVector3(16, 14, 16),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1.0, 0),
		45,
		int32(rl.CAMERA_CUSTOM),
	)

	image := rl.LoadImage("../models/resources/cubicmap.png")
	cubicmap := rl.LoadTextureFromImage(image)
	defer rl.UnloadTexture(cubicmap)

	mesh := rl.GenMeshCubicmap(image, rl.NewVector3(1, 1, 1))

	model := rl.LoadModelFromMesh(mesh)
	defer rl.UnloadModel(model)

	texture := rl.LoadTexture("../models/resources/cubicmap_atlas.png")
	defer rl.UnloadTexture(texture)

	model.Materialser(0).Mapser(rl.MAP_DIFFUSE).Texture = rl.Texture(texture)

	rl.UnloadImage(image)
	rl.SetCameraMode(camera, int32(rl.CAMERA_ORBITAL))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(rl.Camera3D(camera))

		rl.DrawModel(model, rl.NewVector3(-16, 0, -8), 1.0, rl.White)

		rl.EndMode3D()

		rl.DrawTextureEx(cubicmap,
			rl.NewVector2(float32(screenWidth-cubicmap.Width*4-20), 20),
			0, 4.0,
			rl.White,
		)
		rl.DrawRectangleLines(screenWidth-cubicmap.Width*4-20, 20, cubicmap.Width*4, cubicmap.Height*4, rl.Green)

		rl.DrawText("cubicmap image used to", 658, 90, 10, rl.Gray)
		rl.DrawText("generate map 3d model", 658, 104, 10, rl.Gray)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
