package basic

import (
	"fmt"
)

var a, b bool

//TestVariables main function shows variables in go
func TestVariables() {
	{
		fmt.Printf("-----Variables-----\n")

		var c int
		fmt.Println(a, b, c)

		var d, e int = 5, 6
		var f, g, h = false, "string", 5.312 //without type because initializer exists
		fmt.Println(d, e, f, g, h)

		i := 321 //available only inside a function, implicit conversion
		fmt.Println(i)
	}
	{
		var (
			boool          = false //known type so ommited bool
			maxUint uint64 = 1<<64 - 1
			comp           = -5 + 12i //known type so ommited complex128
		)
		fmt.Printf("Type: %T, Value %v\n", boool, boool)
		fmt.Printf("Type: %T, Value %v\n", maxUint, maxUint)
		fmt.Printf("Type: %T, Value %v\n", comp, comp)

		//other types
		//bool string
		//int int8 int16 int32 int64
		//uint uint8 uint16 uint32 uint64 uintptr
		//byte alias for uint8
		//rune alias for int32, represents a Unicode code point
		//float32, float64
		//complex64, complex128
		//uint, int, uintptr depends on architecture
	}

	{
		var c int
		var d float64
		var e bool
		var f string
		fmt.Printf("%v %v %v %q\n", c, d, e, f)
	}

	{
		i := 55
		f := float64(i)
		fmt.Printf("%v, type: %T\n", f, f)
		z := 5.0
		fmt.Printf("%v, type: %T\n", z, z)
	}

	{
		const stringVariable = "const string"
		fmt.Println(stringVariable)

		const (
			big   = 1 << 44
			small = 1 >> 10
		)
		fmt.Print(big, small)
	}
}
