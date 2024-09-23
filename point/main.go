package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func strToRune(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

func setPoint(ptr *point) {
	ptr.x = 24
	ptr.y = 12
}

func convInt(num int) {
	for i := num; i > 0; i /= 10 {
		x := i % 10
		c := '0'
		for j := 0; j < x; j++ {
			c++
		}
		z01.PrintRune(c)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	strToRune("x = ")
	convInt(points.x)
	strToRune(", y = ")
	convInt(points.y)
	z01.PrintRune('\n')
}
