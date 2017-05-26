/*
For is Go's "while"
At that point you can drop the semicolons: C's while is spelled for in Go.
*/
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// If you omit the loop condition it loops forever, 
// so an infinite loop is compactly expressed.
func forever() {
	for {

	}
}

