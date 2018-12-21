package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/liftedkilt/farkle/farkle"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	farkle := farkle.NewGame()

	farkle.Play()

	fmt.Println("Game over with final score of:", farkle.Score)
}
