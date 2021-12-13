package main

import "fmt"

type DataListener struct {
	Name string
}

func (d *DataListener) OnUpdate(data string) {
	fmt.Printf("%s got data change: %s\n", d.Name, data)
}
