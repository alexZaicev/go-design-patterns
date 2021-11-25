package main

import "fmt"

func main() {
	var bldr = NewNotificationBuilder()

	bldr.SetTitle("Title")
	bldr.SetSubtitle("SubTitle")
	bldr.SetIcon("Icon")
	bldr.SetImage("Image")
	bldr.SetMessage("Message")
	bldr.SetPriority(1)
	bldr.SetNotType("NotType")

	n, err := bldr.Build()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Printf("%+v\n", n)
	}

}
