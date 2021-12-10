package main

import "fmt"

type IPod struct {
	IsOn     bool
	Volume   int
	IsPaused bool
}

func (p *IPod) TurnOn() {
	fmt.Println("Turning ipod on")
	p.IsOn = true
}
func (p *IPod) TurnOff() {
	fmt.Println("Turning ipod off")
	p.IsOn = false
}

func (p *IPod) IncreaseVolume() int {
	fmt.Println("Increasing ipod volume")
	p.Volume += 1
	return p.Volume
}
func (p *IPod) DecreaseVolume() int {
	fmt.Println("Decreasing ipod volume")
	p.Volume -= 1
	return p.Volume
}
func (p *IPod) Pause() {
	fmt.Println("Pausing current song")
	p.IsPaused = true
}

func (p *IPod) Play() {
	fmt.Println("Playing current song")
	p.IsPaused = false
}
