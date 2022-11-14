// This is based off a 17 minute tutorial by Anastasia Marchenkova;
// https://www.youtube.com/watch?v=Yo2xmii7TbA

// package is for reusable pieces of code. "Main" is one of such packages. it makes the program executable
// main tells the compiler to process the package as an executable not as a shared library
package main

import (
	// fmt is a Go package that is used to format basic strings, values, inputs, and outputs.
	// without 'fmt', strings don't print!!!
	"fmt"
	//
)

func main() {
	fmt.Println("Hello, World!")
}

// i can build this type-scriptly, in that everytype is empahsized, or I can build it freely. I choose type
// initiated a variable, tF, that is a bool. when ran, I should get a True or False val
var tF bool

// initiated a string, stringy, that is a
var stringy string
