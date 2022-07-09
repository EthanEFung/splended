package main

import (
	"fmt"

	"github.com/ethanefung/splended/models"
)

func main() {
	game := models.NewGame()
	fmt.Println(game.String())
}