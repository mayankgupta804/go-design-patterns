package main

type WalkmanAdapter struct {
	w *Walkman
}

func (wAdap *WalkmanAdapter) IncreaseVolume() int {
	return wAdap.w.VolumeUp()
}

func (wAdap *WalkmanAdapter) DecreaseVolume() int {
	return wAdap.w.VolumeDown()
}

func (wAdap *WalkmanAdapter) TurnOn() {
	wAdap.w.TurnOn()
}

func (wAdap *WalkmanAdapter) TurnOff() {
	wAdap.w.TurnOff()
}

func (wAdap *WalkmanAdapter) Pause() {
	wAdap.w.Pause()
}

func (wAdap *WalkmanAdapter) Play() {
	wAdap.w.Play()
}
