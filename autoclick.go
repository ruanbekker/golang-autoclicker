package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func moveAndClickStepOne() {
	fmt.Println("Moving Mouse")
	robotgo.MoveMouseSmooth(166, 775, 1.0, 7.0)
	fmt.Println("Clicking")
	robotgo.Click()
}

func moveAndClickStepTwo() {
	fmt.Println("Moving Mouse")
	robotgo.MoveMouseSmooth(268, 112, 1.0, 7.0)
	fmt.Println("Clicking")
	robotgo.Click()
}

func delay() {
	fmt.Println("Delaying with 2 Seconds")
	time.Sleep(2 * time.Second)
}

func main() {
	count := 0
	iterations := 4

	for count <= iterations {
		fmt.Println("Starting iteration")
		moveAndClickStepOne()
		delay()
		moveAndClickStepTwo()
		delay()
		fmt.Println("finising iteration")
		count++
	}
	fmt.Println("finished all iterations")
}
