package main

import "fmt"

func main() {
	lib.IterateBooks(MyBookCallback)

	iter := lib.CreateIterator()
	for iter.HasNext() {
		fmt.Printf("Book type: %s\n", iter.Next().Author)
	}
}

func MyBookCallback(b Book) error {
	fmt.Printf("Book title: %s\n", b.Name)
	return nil
}
