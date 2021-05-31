
// +heroku goVersion go1.16


module github.com/nonunknown/kobu-heroku

go 1.16

require (
	github.com/chunqian/go-raylib v0.0.0-20210526135805-f5762ada9240
	github.com/gorilla/websocket v1.4.2
)

replace github.com/chunqian/go-raylib v0.0.0-20210526135805-f5762ada9240 => ./local/github.com/chunqian/go-raylib
replace github.com/gorilla/websocket v1.4.2 => ./local/github.com/gorilla/websocket