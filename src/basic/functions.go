package basic

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func multiplyAndDevide(x, y int) (int, float64) {
	return x * y, float64(x) / float64(y)
}

func addAndSubtract(x, y int) (sum, difference int) { //sum and difference declaration
	sum = x + y
	difference = x - y
	return
}

//TestFunctions main function in basic1
func TestFunctions() {
	fmt.Printf("-----Functions-----\nSum = %d\n", add(4, 5))

	a, b := multiplyAndDevide(4, 5)
	fmt.Printf("Product = %d, quotient = %f\n", a, b)

	c, d := addAndSubtract(8, 21)
	fmt.Printf("Sum = %d, difference = %d\n", c, d)
}
