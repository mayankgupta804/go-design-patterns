package main

type MusicPlayer interface {
	IncreaseVolume() int
	DecreaseVolume() int
	TurnOn()
	TurnOff()
	Pause()
	Play()
}


