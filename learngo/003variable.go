package learngo

import (
	"fmt"
)

//variable

//number starting

//go type system

// constant
func integerS() {
	fmt.Println(
		-200, -300, 0, 50, 100,
	)
}

func floatS() {
	fmt.Println(
		.2, -0., 1., 0.0, 10.56,
	)
}

func booleaN() {
	fmt.Println(true, false)
}

func strinG() {
	fmt.Println(
		"",            //empty string
		"hi",          // string
		"क्या हाल है", // How are you in hindi
	)
}

func hexadecimaL() {
	fmt.Println(0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)
	fmt.Println(0xa, 0xb, 0xc, 0xd, 0xe, 0xf)
	fmt.Println(0x11) // 17
	fmt.Println(0x19) // 25
	fmt.Println(0x32) // 50
	fmt.Println(0x64) // 100
}

func basicDatatypes() {
	integerS()
	floatS()
	booleaN()
	strinG()
	hexadecimaL()
}

func varDeclaration() {
	// integer types
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// float types
	var f32 float32
	var f64 float64

	// complex types
	var c64 complex64
	var c128 complex128

	// bool type
	var b bool

	// string types
	var s string
	var r rune  // also a numeric type
	var by byte // also a numeric type
	var sli []string

	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)

	// You could do it with Println as well
	fmt.Printf("%q\n", s)

	fmt.Println(sli)

}

// VariableDatatypes prints all Datatypes
func VariableDatatypes() {
	basicDatatypes()
	varDeclaration()

	fmt.Printf("%08b\n", 0)
	fmt.Printf("%08b\n", 1)
	fmt.Printf("%08b\n", 2)
	fmt.Printf("%08b\n", 3)
	fmt.Printf("%08b\n", 4)
	fmt.Printf("%08b\n", 5)
	fmt.Printf("%08b\n", 6)
	fmt.Printf("%08b\n", 7)
	fmt.Printf("%08b\n", 8)
}
