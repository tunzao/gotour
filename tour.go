package main

// factored import statement
import (
	"fmt"
	"math"
	"math/rand"
)

// add A function can take zero or more arguments
func add(x int, y int) int {
	return x + y
}

// swap a function can return any numbers of result
func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// name starts with a capital letter is exported
	fmt.Println("Pi=", math.Pi)
}
