package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	up := "FFED"
	down := "FFEC"
	right := "FFEA"
	left := "FFEB"
	// up := keyboard.Key.KeyArrowUp
	// down := keyboard.Key.KeyArrowDown
	// right := keyboard.Key.KeyArrowRight
	// left := keyboard.Key.KeyArrowLeft
	var angleConst float64 = 10
	var t turtle = turtle{0, 0}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Game started. Press ESC to quit")
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch fmt.Sprintf("%X", key) {
		case up:
			t.forward()
		case down:
			t.back()
		case right:
			t.rotate(angleConst)
		case left:
			t.rotate(-angleConst)
		}

		fmt.Printf("You pressed key %X\r\n", key)
		t.print()
		if key == keyboard.KeyEsc {
			break
		}
	}
}
