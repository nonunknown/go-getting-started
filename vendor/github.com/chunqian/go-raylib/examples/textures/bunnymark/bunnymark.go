package main

import (
	rl "github.com/chunqian/go-raylib/raylib"

	"fmt"
	"runtime"
)

var (
	MAX_BUNNIES        = 50000
	MAX_BATCH_ELEMENTS = 8192
)

type Bunny struct {
	position rl.Vector2
	speed    rl.Vector2
	color    rl.Color
}

func init() {
	runtime.LockOSThread()
}

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [textures] example - bunnymark")
	defer rl.CloseWindow()

	texBunny := rl.LoadTexture("../textures/resources/wabbit_alpha.png")
	defer rl.UnloadTexture(texBunny)

	bunnies := make([]Bunny, MAX_BUNNIES)
	bunniesCount := 0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsMouseButtonDown(int32(rl.MOUSE_LEFT_BUTTON)) {

			for i := 0; i < 100; i++ {

				if bunniesCount < MAX_BUNNIES {
					bunnies[bunniesCount].position = rl.GetMousePosition()

					bunnies[bunniesCount].speed.X = float32(rl.GetRandomValue(-250, 250)) / 60.0
					bunnies[bunniesCount].speed.Y = float32(rl.GetRandomValue(-250, 250)) / 60.0

					bunnies[bunniesCount].color = rl.NewColor(
						byte(rl.GetRandomValue(50, 240)),
						byte(rl.GetRandomValue(80, 240)),
						byte(rl.GetRandomValue(100, 240)),
						255,
					)
					bunniesCount++
				}
			}
		}

		for i := 0; i < bunniesCount; i++ {

			bunnies[i].position.X += bunnies[i].speed.X
			bunnies[i].position.Y += bunnies[i].speed.Y

			if (bunnies[i].position.X+float32(texBunny.Width)/2.0) > float32(rl.GetScreenWidth()) ||
				(bunnies[i].position.X+float32(texBunny.Width)/2.0) < 0.0 {
				bunnies[i].speed.X *= -1.0
			}
			if (bunnies[i].position.Y+float32(texBunny.Height)/2.0) > float32(rl.GetScreenHeight()) ||
				(bunnies[i].position.Y+float32(texBunny.Height)/2.0-40.0) < 0.0 {
				bunnies[i].speed.Y *= -1.0
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for i := 0; i < bunniesCount; i++ {
			rl.DrawTexture(texBunny, int32(bunnies[i].position.X), int32(bunnies[i].position.Y), bunnies[i].color)
		}

		rl.DrawRectangle(0, 0, screenWidth, 40, rl.Black)

		rl.DrawText(fmt.Sprintf("bunnies: %d", bunniesCount), 120, 10, 20, rl.Green)

		rl.DrawText(fmt.Sprintf("batched draw calls: %d", 1+bunniesCount/MAX_BATCH_ELEMENTS), 320, 10, 20, rl.Maroon)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
