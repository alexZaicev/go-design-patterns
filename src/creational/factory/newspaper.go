package main

import "fmt"

type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("This is newspaper name %s", n.name)
}

func createNewspaper(name string, pages int, publisher string) IPublication  {
	return &newspaper{
		publication: publication{
			name: name,
			pages: pages,
			publisher: publisher,
		},
	}
}
