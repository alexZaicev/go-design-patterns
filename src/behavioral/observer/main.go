package main

import "fmt"

func main() {

	listener1 := &DataListener{
		Name: "Listener 1",
	}

	listener2 := &DataListener{
		Name: "Listener 2",
	}

	subject := &DataSubject{}

	fmt.Println("########################################")
	subject.registerObserver(listener1)
	subject.registerObserver(listener2)
	subject.ChangeItem("data")

	fmt.Println("########################################")
	subject.unregisterObserver(listener2)
	subject.ChangeItem("data")
}
