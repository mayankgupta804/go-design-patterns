package main

import "fmt"

type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("This newspaper is named: %s", n.name)
}

func createNewspaper(name string, pub string, pg int) (IPublication, error) {
	return &newspaper{publication{
		name:      name,
		pages:     pg,
		publisher: pub,
	}}, nil
}
