package main

import "fmt"

func main() {
	magazine, _ := newPublication("magazine", "Mayank", "New York Times", 10)
	newspaper, _ := newPublication("newspaper", "Gupta", "Times of India", 10)

	fmt.Println(magazine)
	fmt.Println(newspaper)
}
