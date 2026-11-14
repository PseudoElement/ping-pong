package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pseudoelement/ping-pong/src/game"
)

type FontLoader struct {
	fonts []rl.Font
}

func main() {
	game := game.NewGame()
	game.Start()
}
