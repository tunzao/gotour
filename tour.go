package main

// factored import statement
import (
	"fmt"
	"math"
	"math/rand"
)

const (
	// an untyped constant takes the type needed in the context
	Big   = 1 << 100
	Small = 1 >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

// package level variables with initializers
var c, python, java = true, false, "No!"

// add A function can take zero or more arguments
func add(x int, y int) int {
	return x + y
}

// swap a function can return any numbers of result
func swap(x, y int) (int, int) {
	return y, x
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// '"naked" return' returns the named values
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// name starts with a capital letter is exported
	fmt.Println("Pi=", math.Pi)

	// function level variable
	// variables declared without explicit init values are given their zero value
	// 0 for numeric types
	// false for boolean type
	// "" for strings
	var i int

	// := short assignment only works in a function
	k := 2
	fmt.Println(i, c, python, java, k)

	v := 20
	fmt.Printf("v is of type %T\n", v)

	// declare constants with const keyword, const cannot be declared by short assignment :=
	const running = 1

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
