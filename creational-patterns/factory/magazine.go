package main

import "fmt"

type magazine struct {
	publication
}

func (n magazine) String() string {
	return fmt.Sprintf("This magazine is named: %s", n.name)
}

func createMagazine(name string, pub string, pg int) (IPublication, error) {
	return &magazine{publication{
		name:      name,
		pages:     pg,
		publisher: pub,
	}}, nil
}
