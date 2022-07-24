package main

// factored import statement
import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
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

func Sqrt(x float64) float64 {
	z, loop := x/2, 0

	for pre := 0.0; math.Abs(z-pre) >= 0.00001; loop++ {
		pre = z
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println("loop taken", loop)
	return z
}

func osinfo() {
	fmt.Print("go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macos.")
		// break statement is automatically provided by GO
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

// the defer statement defers the execution of a function until the surround function returns.
// what's surround function?!
func deferState() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done.")

}

func pointerDemo() {
	i, j := 12, 34

	p := &i
	fmt.Println(*p)
	*p = 100
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	*p = *p / 10
	fmt.Println(j)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			pic[i][j] = uint8(i * j)
		}
	}
	return pic
}

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	wc := make(map[string]int)
	for _, field := range fields {
		value, _ := wc[field]
		value++
		wc[field] = value
	}
	return wc
}

func fibonacci() func() int {
	x, y, n := 0, 1, 0
	return func() int {
		n++
		switch n - 1 {
		case 0:
			return 0
		case 1:
			return 1
		default:
			r := x + y
			x, y = y, r
			return r
		}
	}
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
	fmt.Println()

	println(math.Sqrt(4) - Sqrt(4))
	println(math.Sqrt(2) - Sqrt(2))
	fmt.Println()

	osinfo()
	fmt.Println()

	deferState()
	fmt.Println()

	pointerDemo()
	fmt.Println()

	fmt.Println(WordCount("hello world is everything ok ok"))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
