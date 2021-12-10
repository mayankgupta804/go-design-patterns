package main

import "errors"

func newPublication(pubType, pub, name string, pg int) (IPublication, error) {
	switch pubType {
	case "newspaper":
		return createNewspaper(name, pub, pg)
	case "magazine":
		return createMagazine(name, pub, pg)
	default:
		return nil, errors.New("invalid publication type")
	}
}


