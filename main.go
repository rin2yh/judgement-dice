package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"judgement-dice/internal/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("ジャッジメントダイス")
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatal(err)
	}
}
