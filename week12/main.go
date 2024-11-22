package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Craet a slice by make function
	var gpaSlice []float64
	gpaSlice[0] = 4.1
	gpaSlice[1] = 4.5
	gpaSlice[2] = 3.9
	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// Create a slice by slice literal
	// gpaSlice := []float64{4.1,4.5,3.9}  // slice literal
	// fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// Create a slice by slicing an existing array
	// gpas := [5]float64{3.5, 4.1, 4.5, 4.9, 4.23} // array := array literal
	// fmt.Print(gpas, reflect.TypeOf(gpas))
	// gpaSlice := gpas[1:3] // slice := slicing array
	// fmt.Print(gpaSlice, reflect.TypeOf(gpaSlice))
}
