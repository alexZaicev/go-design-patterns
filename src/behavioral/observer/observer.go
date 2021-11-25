package main

import "fmt"

type observer interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

func (listener *DataListener) onUpdate(data string) {
	fmt.Println("Listener:", listener.Name, "got data change:", data)
}
