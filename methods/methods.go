// In Go, a method is a function that is associated with a particular type (usually a struct). It allows you to define behavior that operates on instances of that type. Methods can have parameters and return values just like regular functions, but they also have a special receiver argument that specifies which type the method is associated with.

package main

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type Triangle struct {
	side1, side2, side3 float64
	base, height        float64
}

func (t Triangle) area() {
	fmt.Println(0.5 * t.base * t.height)
}

func (t Triangle) perimeter() float64 {
	return t.side1 + t.side2 + t.side3
}

func main() {
	var choice int
	fmt.Println("Geometry Calculator")
	fmt.Println("Choose a shape to calculate:")
	fmt.Println("1. Rectangle")
	fmt.Println("2. Circle")
	fmt.Println("3. Triangle")
	fmt.Print("Enter your choice (1-3): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var width, height float64
		fmt.Println("Enter width of Rectangle: ")
		fmt.Scanln(&width)
		fmt.Println("Enter height of Rectangle: ")
		fmt.Scanln(&height)

		rect := Rectangle{width: width, height: height}
		fmt.Println("Area of the rectangle : ", rect.area())
		fmt.Println("Perimeter of the rectangle : ", rect.perimeter())

	case 2:
		var radius float64
		fmt.Println("Enter the radius of the circle: ")
		fmt.Scanln(&radius)

		circle := Circle{radius: radius}
		fmt.Println("Area of the rectangle : ", circle.area())
		fmt.Println("Perimeter of the rectangle : ", circle.perimeter())

	case 3:
		var side1, side2, side3, base, height float64
		fmt.Print("Enter base of Triangle: ")
		fmt.Scan(&base)
		fmt.Print("Enter height of Triangle: ")
		fmt.Scan(&height)
		fmt.Print("Enter lengths of sides (side1, side2, side3): ")
		fmt.Scan(&side1, &side2, &side3)

		tr := Triangle{side1: side1, side2: side2, side3: side3, base: base, height: height}
		fmt.Println("Area of the rectangle : ")
		tr.area()
		fmt.Println("Perimeter of the rectangle : ", tr.perimeter())

	}
}
