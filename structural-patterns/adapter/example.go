package main

func main() {
	var ipod MusicPlayer = &IPod{}
	var walkman MusicPlayer = &WalkmanAdapter{w: &Walkman{}}

	ipod.TurnOn()
	ipod.IncreaseVolume()
	ipod.Play()

	walkman.TurnOn()
	walkman.Play()
	walkman.Pause()
}
