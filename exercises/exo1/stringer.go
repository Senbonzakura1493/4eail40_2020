package main

import (
	"fmt"
	"math"
)

// Implement types Rectangle, Circle and Shape
type Shape interface {
	Area() float64
}
type Circle struct {
	r float64
}

type Rectangle struct {
	width  float64
	length float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length

}
func (c Circle) Area() float64 {
	return 3.14 * math.Pow(c.r, 2)
}
func (r Rectangle) String() string {
	return fmt.Sprintf("Square of width %f and length %f", r.width, r.length)
}
func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %f", c.r)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
