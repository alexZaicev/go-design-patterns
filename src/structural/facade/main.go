package main

import "fmt"

func main() {
	fmt.Println("Making americano...")
	fmt.Println("########################################")
	makeAmericano(8.0)
	fmt.Println("########################################")
	fmt.Println("Making cafe latte...")
	fmt.Println("########################################")
	makeLatte(12.0, true)
}
