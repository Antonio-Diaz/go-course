//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
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

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate
	b Coordinate
}

func width(r Rectangle) int {
	return r.b.x - r.a.x
}

func length(r Rectangle) int {
	return r.a.y - r.b.y
}

func area(r Rectangle) int {
	return length(r) * width(r)
}

func perimeter(r Rectangle) int {
	return 2 * (width(r) + length(r))
}

func printInfo(r Rectangle) {
	println("Area:", area(r))
	println("Perimeter:", perimeter(r))
}

func main() {
	rectangle := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rectangle)

	rectangle.a.y *= 2
	rectangle.b.x *= 2

	printInfo(rectangle)
}
