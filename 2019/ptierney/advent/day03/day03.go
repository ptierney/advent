package day03

import (
	"advent/common"

	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	p := new(Point)
	p.X = x
	p.Y = y
	return p
}

func getOffsetPoint(startPt *Point, cmd string) *Point {
	direction := string(cmd[0])

	amountStr := strings.TrimLeft(cmd, "URDL")

	amount, err := strconv.Atoi(amountStr)

	if err != nil {
		panic(err)
	}

	var x int = 0
	var y int = 0

	switch direction {
	case "U":
		y = amount
	case "D":
		y = -amount
	case "R":
		x = amount
	case "L":
		x = -amount
	}

	return NewPoint(startPt.X+x, startPt.Y+y)
}

func createWire(input string) []*Point {
	wire := make([]*Point, 1)

	wire[0] = NewPoint(0, 0)

	wireCmds := strings.Split(input, ",")

	for i, cmd := range wireCmds {
		lastPt := wire[i]

		pt := getOffsetPoint(lastPt, cmd)

		wire = append(wire, pt)
	}

	return wire
}

type Orientation int

const (
	Vertical   Orientation = iota
	Horizontal Orientation = iota
)

func segmentsIntersect(w1_s, w1_e, w2_s, w2_e *Point) []*Point {
	var wire1Otn Orientation
	var wire2Otn Orientation

	if w1_s.Y == w1_e.Y {
		wire1Otn = Horizontal
	} else {
		wire1Otn = Vertical
	}

	if w2_s.Y == w2_e.Y {
		wire2Otn = Horizontal
	} else {
		wire2Otn = Vertical
	}

	intersectionPoints := make([]*Point, 0)

	if wire1Otn == Horizontal && wire2Otn == Vertical {
		// Wire 1 is horizontal, wire 2 is vertical
		wire1MaxX := common.Max(w1_s.X, w1_e.X)
		wire1MinX := common.Min(w1_s.X, w1_e.X)

		wire2MaxY := common.Max(w2_s.Y, w2_e.Y)
		wire2MinY := common.Min(w2_s.Y, w2_e.Y)

		intersects := ((w2_s.X > wire1MinX) && (w2_s.X < wire1MaxX)) &&
			((w1_s.Y > wire2MinY) && (w1_s.Y < wire2MaxY))

		if intersects == false {
			return intersectionPoints
		}

		pt := NewPoint(w2_s.X, w1_s.Y)

		intersectionPoints = append(intersectionPoints, pt)

		return intersectionPoints
	} else if wire1Otn == Vertical && wire2Otn == Horizontal {
		// Wire 1 is vertical, wire 2 is horizontal
		wire2MaxX := common.Max(w2_s.X, w2_e.X)
		wire2MinX := common.Min(w2_s.X, w2_e.X)

		wire1MaxY := common.Max(w1_s.Y, w1_e.Y)
		wire1MinY := common.Min(w1_s.Y, w1_e.Y)

		intersects := ((w1_s.X > wire2MinX) && (w1_s.X < wire2MaxX)) &&
			((w2_s.Y > wire1MinY) && (w2_s.Y < wire1MaxY))

		if intersects == false {
			return intersectionPoints
		}

		pt := NewPoint(w1_s.X, w2_s.Y)

		intersectionPoints = append(intersectionPoints, pt)

		return intersectionPoints
	} else if wire1Otn == Vertical && wire2Otn == Vertical {
		// Both wires are vertical
		return intersectionPoints
	} else if wire1Otn == Horizontal && wire2Otn == Horizontal {
		// Both wires are horizontal
		return intersectionPoints
	}

	panic("Should not reach here")
}

func getIntersections(wire1, wire2 []*Point) []*Point {
	intersections := make([]*Point, 0)

	for i := 0; i < len(wire1)-1; i++ {
		for j := 0; j < len(wire2)-1; j++ {
			pts := segmentsIntersect(wire1[i], wire1[i+1],
				wire2[j], wire2[j+1])

			if len(pts) == 0 {
				continue
			}

			intersections = append(intersections, pts...)
		}
	}

	return intersections
}

func pointIntersectsSegment(pt, seg_s, seg_e *Point) bool {
	var wireOtn Orientation

	if seg_s.Y == seg_e.Y {
		wireOtn = Horizontal
	} else {
		wireOtn = Vertical
	}

	if wireOtn == Horizontal {
		maxX := common.Max(seg_s.X, seg_e.X)
		minX := common.Min(seg_s.X, seg_e.X)

		return (pt.X >= minX && pt.X <= maxX) && (pt.Y == seg_s.Y)
	} else {
		maxY := common.Max(seg_s.Y, seg_e.Y)
		minY := common.Min(seg_s.Y, seg_e.Y)

		return (pt.Y >= minY && pt.Y <= maxY) && (pt.X == seg_s.X)
	}

	panic("should not be here")
	return false
}

func segmentLength(sp, ep *Point) int {
	return common.Abs(sp.X-ep.X) + common.Abs(sp.Y-ep.Y)
}

func getWireDistanceToIntersection(wire []*Point, intersection *Point) int {
	var sum int = 0

	for i := 0; i < len(wire)-1; i++ {
		if pointIntersectsSegment(intersection, wire[i], wire[i+1]) == false {
			sum += segmentLength(wire[i], wire[i+1])
		} else {
			sum += segmentLength(wire[i], intersection)
			return sum
		}
	}

	return sum
}

func SolvePart2(input []string) int {
	wire1 := createWire(input[0])
	wire2 := createWire(input[1])

	intersections := getIntersections(wire1, wire2)

	var smallestDist int = 0

	for i, inter := range intersections {
		wire1Dist := getWireDistanceToIntersection(wire1, inter)
		wire2Dist := getWireDistanceToIntersection(wire2, inter)

		totalDist := wire1Dist + wire2Dist

		if i == 0 {
			smallestDist = totalDist
			continue
		}

		if totalDist > smallestDist {
			continue
		}

		smallestDist = totalDist
	}

	return smallestDist
}

func SolveProblem(input []string) int {
	wire1 := createWire(input[0])
	wire2 := createWire(input[1])

	intersections := getIntersections(wire1, wire2)

	var smallestDist int

	for i, isect := range intersections {
		dist := common.Abs(isect.X) + common.Abs(isect.Y)

		if i == 0 {
			smallestDist = dist
			continue
		}

		if dist < smallestDist {
			smallestDist = dist
		}
	}

	return smallestDist
}
