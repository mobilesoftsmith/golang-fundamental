package main

import (
	"errors"
	"log"
	"fmt"
)

func main() {
	fmt.Println("error:")
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
		// return 0, fmt.Errorf("norgate math: square root of negative number")
		// return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
	}
	// implementation
	return 42, nil
}
