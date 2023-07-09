package main

import (
	"codes/task1"
	"codes/task2"
)

func main() {
	task1.TrackMatch("task1/input.json")
	task2.GetRoadTrafficStatus("珞喻路","武汉")
	task2.RoutePlanning()
}
