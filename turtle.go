package main

import (
	"fmt"
	"math"
)

type turtle struct {
	angle    float64
	distance float64 // distance along the angle
}

func (t *turtle) back() {
	t.distance -= 1
}
func (t *turtle) forward() {
	t.distance += 1
}

func (t *turtle) rotate(theta float64) {
	t.angle += theta
}

func (t *turtle) polar2cartesian() [2]float64 {
	x := t.distance * math.Cos(t.angle)
	y := t.distance * math.Sin(t.angle)
	out := [2]float64{x, y}
	return out
}
func (t *turtle) print() {
	x := t.polar2cartesian()
	fmt.Printf("x,y location is: (%f, %f) . polar distance %f angle %f", x[0], x[1], t.distance, t.angle)
}

// func main() {
// 	var t turtle = turtle{0, 0}
// 	cmds := []int{10, 10, 30}
// 	for _, e := range cmds {
// 		t.move()
// 		t.rotate(float64(e))
// 		t.print()

// 	}
// }
