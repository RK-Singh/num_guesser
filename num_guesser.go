package main

import (
	"fmt"
	"time"
)

func printString(s string) {
	fmt.Println(s)
	for i := 0; i < 40; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Print("-")
	}
	fmt.Println("100%")
}

// write a function to check if a number is a prime number
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	var err error
	var progress_messages = []string{"Analysing brainwaves...", "Scanning memories...", "Calculating probabilities...", "Decoding thoughts..."}

	fmt.Println("Guess a number between 1 and 10:")
	_, err = fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Println("Invalid input, integer value expected.")
	} else if number < 1 || number > 10 {
		fmt.Println("Invalid input, number must be between 1 and 10.")
	} else {

		for _, message := range progress_messages {
			printString(message)
		}
		fmt.Printf("You guessed the number %d.\n", number)
	}
}
