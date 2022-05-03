package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("FizzBuzz\n")
		}
		if i%3 == 0 {
			fmt.Printf("Fizz\n")
		}
		if i%5 == 0 {
			fmt.Printf("Buzz\n")
		}
		if i != 0 {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}

// Run
// go run fizzbuzz.go