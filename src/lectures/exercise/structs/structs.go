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

import "fmt"

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func scaleShape(rectangle Rectangle, scale float64) Rectangle {
	return Rectangle{
		x1: rectangle.x1,
		x2: (rectangle.x1 + scale*(rectangle.x2-rectangle.x1)),
		y1: rectangle.y1,
		y2: (rectangle.y1 + scale*(rectangle.y2-rectangle.y1))}
}

func getArea(rectangle Rectangle) float64 {
	return (rectangle.x2 - rectangle.x1) * (rectangle.y2 - rectangle.y1)
}

func getPerimeter(rectangle Rectangle) float64 {
	return 2*(rectangle.x2-rectangle.x1) + 2*(rectangle.y2-rectangle.y1)
}

func analyze(rectangle Rectangle) {
	fmt.Println(rectangle)
	fmt.Println("with area", getArea(rectangle))
	fmt.Println("and with perimeter", getPerimeter(rectangle))
}

func main() {
	rectangle := Rectangle{1.0, 3.0, 0.0, 2.0}
	analyze(rectangle)
	analyze(scaleShape(rectangle, 2))
}
