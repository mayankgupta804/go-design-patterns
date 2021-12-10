package main

import "fmt"

type Walkman struct {
	IsOn     bool
	Volume   int
	IsPaused bool
}

func (w *Walkman) TurnOn() {
	fmt.Println("Turning walkman on")
	w.IsOn = true
}

func (w *Walkman) TurnOff() {
	fmt.Println("Turning walkman off")
	w.IsOn = false
}

func (w *Walkman) VolumeUp() int {
	fmt.Println("Increasing walkman volume")
	w.Volume += 1
	return w.Volume
}

func (w *Walkman) VolumeDown() int {
	fmt.Println("Decreasing walkman volume")
	w.Volume -= 1
	return w.Volume
}
func (w *Walkman) Pause() {
	fmt.Println("Pausing current song")
	w.IsPaused = true
}

func (w *Walkman) Play() {
	fmt.Println("Playing current song")
	w.IsPaused = false
}
