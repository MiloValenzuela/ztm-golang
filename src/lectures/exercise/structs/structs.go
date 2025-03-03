//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

//* Create a rectangle structure containing a length and width field
type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}

func length(rect Rectangle) int {
	return (rect.a.y - rect.a.x)
}

//* Using functions, calculate the area and perimeter of a rectangle,
func area(rect Rectangle) int {
	return length(rect) * width(rect)
}

func perimeter(rect Rectangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

// - Print the results to the terminal
func printInfo(rect Rectangle) {
	fmt.Println("area is", area(rect))
	fmt.Println("Perimeter is", perimeter((rect)))
}

//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations

func main() {
	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	//  - Print the new results to the terminal
	printInfo(rect)

	rect.a.y *= 2
	rect.a.x *= 2

	printInfo(rect)
}
