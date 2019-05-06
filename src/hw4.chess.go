package main

import "fmt"

// Point to calculate the chess horse
type Point struct {
	x, y int
}

func (point Point) nextPoint(dx, dy int) (next Point, ok bool) {
	next = Point{point.x + dx, point.y + dy}
	ok = next.x > -1 && next.y > -1 && next.x < 8 && next.y < 8
	return
}

func (point Point) nextPoints() []Point {
	points := make([]Point, 0, 8)
	dxy := [4][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}}
	for _, xy := range dxy {
		for _, k := range [2]int{1, -1} {
			if next, ok := point.nextPoint(xy[0], xy[1]*k); ok {
				points = append(points, next)
			}
		}
	}
	return points
}

func main() {
	point2x2 := Point{2, 2}
	point0x2 := Point{0, 2}
	point0x0 := Point{0, 0}
	point7x7 := Point{7, 7}
	fmt.Println(point2x2.nextPoints())
	fmt.Println(point0x2.nextPoints())
	fmt.Println(point0x0.nextPoints())
	fmt.Println(point7x7.nextPoints())
}
