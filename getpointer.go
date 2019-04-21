package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func get() {
	time.Sleep(2 * time.Second)
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
}

func main() {
	count := 0
	max := 20

	for count < max {
		// delay with 2 seconds
		time.Sleep(2 * time.Second)

		// get the mouse coordinates
		get()

		// increment counter
		count++
	}
	fmt.Println("exiting..")
}
