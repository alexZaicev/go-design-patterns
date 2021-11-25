package main

import "fmt"

func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}

func callbackTest()  {
	fmt.Println("########################################")
	fmt.Println("Callback function....")
	lib.IterateBooks(myBookCallback)
	fmt.Println("########################################")
	fmt.Println("Anonymous callback function....")
	lib.IterateBooks(func(b Book) error {
		fmt.Println("Book author:", b.Author)
		return nil
	})

	fmt.Println()
	fmt.Println()
}

func interfaceTest() {
	fmt.Println("########################################")
	it := lib.CreateIterator()
	for it.hasNext() {
		book := it.next()
		fmt.Println("Book title:", book.Name)
	}
}

func main() {
	callbackTest()
	interfaceTest()
}
