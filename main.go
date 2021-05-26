package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	rl "github.com/chunqian/go-raylib/raylib"
	"github.com/gorilla/websocket"
)

// var addr *string;

func init() {
	runtime.LockOSThread()
}

// func test() {
// 	// flag.Parse()
// 	port := os.Getenv("PORT")

// 	if (port == "") {
// 		log.Fatal("Port is not set")

// 	}

// 	log.Println("PORT IS: ", port)
// 	// addr = flag.String("addr", "0.0.0.0:"+string(port), "http service address")
// 	http.HandleFunc("/gameEngine", handler)
// 	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "JumpAndShoot engine is running on this port")
// 	})

// 	err := http.ListenAndServe(*addr, nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

// var text string = "Waiting for godot"

func main() {
	//test
	// flag.Parse()
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("Port is not set")
		port = "5000"

	}

	log.Println("PORT IS: ", port)

	go func() {
		http.HandleFunc("/gameEngine", handler)
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "JumpAndShoot engine is running on this port")
		})

		err := http.ListenAndServe(":"+string(port), nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	// go test()
	// log.Println("Started....")

	// log.Println("Initializing Window ....")

	screenWidth := int32(100)
	screenHeight := int32(100)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.GuiSetState(int32(rl.FLAG_WINDOW_HIDDEN))
	rl.SetTargetFPS(15)
	log.Println("Raylib is running")
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("text", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func handler(w http.ResponseWriter, r *http.Request) {
	playerName := r.URL.Query()["id"][0]
	log.Println("PlayerJoined :", playerName)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ERROR", err)
		return
	}

	conn.WriteMessage(1, []byte("Hello from server"))

	// go func() {
	// 	for {
	// 		_, msg, err := conn.ReadMessage()
	// 		if err != nil {
	// 			log.Println("Player disconnected")
	// 			break
	// 		}
	// 		output := string(msg)
	// 		log.Println(output)
	// 		if output != "" {
	// 			text = output
	// 		}
	// 	}
	// }()

	log.Println("connected: ", conn.RemoteAddr())

	// // Create a new player instance
	// p := objects.NewPlayer(playerName, conn)

	// // Start listening to messages from player
	// go p.RecieveMessages()

	// // Find a match for the player
	// objects.GS.Match(p)
}
