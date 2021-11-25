package main

import "fmt"

const MAGAZINE = "magazine"
const NEWSPAPER = "newspaper"

func NewPublication(pubType string, name string, pages int, publisher string) (IPublication, error){
	switch pubType {
	case MAGAZINE:
		return createMagazine(name, pages, publisher), nil
	case NEWSPAPER:
		return createNewspaper(name, pages, publisher), nil
	}
	return nil, fmt.Errorf("unknown publication type [%s]\n", pubType)
}
