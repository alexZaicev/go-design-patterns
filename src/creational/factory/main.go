package main

import "fmt"

func printDetails(pb IPublication) {
	fmt.Println("##################################")
	fmt.Printf("Name: %s\n", pb.getName())
	fmt.Printf("Pages: %d\n", pb.getPages())
	fmt.Printf("Publisher: %s\n", pb.getPublisher())
	fmt.Println("##################################")
}

func main() {

	var pbArr [4]IPublication
	pbArr[0], _ = NewPublication(MAGAZINE, "Forbes", 50, "Forbes")
	pbArr[1], _ = NewPublication(NEWSPAPER, "New York Times", 50, "New York Times")
	pbArr[2], _ = NewPublication(MAGAZINE, "Economist", 50, "Economist")
	pbArr[3], _ = NewPublication(NEWSPAPER, "Financial Times", 50, "Financial Times")

	for _, pb := range pbArr {
		printDetails(pb)
	}
}
