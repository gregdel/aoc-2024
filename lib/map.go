package aoc

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Direction represents a direction.
type Direction uint8

// Available directions.
const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionRight
	DirectionLeft
	DirectionUpRight
	DirectionDownRight
	DirectionUpLeft
	DirectionDownLeft
	DirectionUnkown
)

func (d Direction) String() string {
	return DirStr(d)
}

var dirStrs = []string{
	"Up", "Down", "Right", "Left",
	"UpRight", "DownRight", "UpLeft", "DownLeft",
	"Unknown",
}

func DirFromRune(r rune) Direction {
	switch r {
	case '^':
		return DirectionUp
	case 'v':
		return DirectionDown
	case '>':
		return DirectionRight
	case '<':
		return DirectionLeft
	default:
		panic("Direction not handled")
	}
}

// DirStr returns the string representation of a direction.
func DirStr(d Direction) string {
	return dirStrs[int(d)]
}

// AllDirection represents all the directions.
var AllDirection = []Direction{
	DirectionDown,
	DirectionLeft,
	DirectionUp,
	DirectionRight,
}

// AllDirectionWithDiags represents all the directions, diagonals included.
var AllDirectionWithDiags = append([]Direction{
	DirectionUpRight,
	DirectionDownRight,
	DirectionUpLeft,
	DirectionDownLeft,
}, AllDirection...)

// OppositeDirection returns the opposite direction of a direction.
func OppositeDirection(d Direction) Direction {
	switch d {
	case DirectionUp:
		return DirectionDown
	case DirectionDown:
		return DirectionUp
	case DirectionRight:
		return DirectionLeft
	case DirectionLeft:
		return DirectionRight
	case DirectionUpRight:
		return DirectionDownLeft
	case DirectionDownRight:
		return DirectionUpLeft
	case DirectionDownLeft:
		return DirectionUpRight
	case DirectionUpLeft:
		return DirectionDownRight
	case DirectionUnkown:
		return DirectionUnkown
	}
	panic("Invalid direction")
}

// RotateCW rotates a direction clockwise
func RotateCW(d Direction) Direction {
	switch d {
	case DirectionUp:
		return DirectionRight
	case DirectionRight:
		return DirectionDown
	case DirectionDown:
		return DirectionLeft
	case DirectionLeft:
		return DirectionUp
	}
	panic("Invalid direction")
}

// Point represents a 2d point.
type Point struct {
	X, Y int
	C    rune
}

// NewPoint returns a new point
func NewPoint(x, y int, c rune) *Point {
	return &Point{X: x, Y: y, C: c}
}

func (p *Point) String() string {
	return fmt.Sprintf("x:%d;y:%d [%c]", p.X, p.Y, p.C)
}

func (p Point) Translate(v Vec) Point {
	p.X += v.U
	p.Y += v.V
	return p
}

func (p Point) Wrap(x, y int) Point {
	p.X %= x
	p.Y %= y
	if p.X < 0 {
		p.X += x
	}
	if p.Y < 0 {
		p.Y += y
	}
	return p
}

// Map2D represents a 2DMap.
type Map2D struct {
	Points [][]*Point
}

// NewMap2D returns a new 2D map.
func NewMap2D() *Map2D {
	return &Map2D{
		Points: [][]*Point{},
	}
}

// NewFixedMap2D returns a new 2D map filled with the given rune.
func NewEmptyMap2D(width, height int, r rune) *Map2D {
	line := strings.Builder{}
	for x := 0; x < width; x++ {
		line.WriteRune(r)
	}

	m := NewMap2D()
	for y := 0; y < height; y++ {
		m.AddPointsFromLine(line.String())
	}

	return m
}

// NewMap2DFromReader returns a new 2D map reading line by line in a reader.
func NewMap2DFromScanner(scanner *bufio.Scanner) *Map2D {
	m := NewMap2D()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return m
		}
		m.AddPointsFromLine(line)
	}
	return m
}

// NewMap2DFromReader returns a new 2D map reading line by line in a reader.
func NewMap2DFromReader(r io.Reader) *Map2D {
	return NewMap2DFromScanner(bufio.NewScanner(r))
}

// AddPointsFromLine adds points to a 2d map from a string.
func (m *Map2D) AddPointsFromLine(line string) {
	y := len(m.Points)
	points := make([]*Point, len(line))
	m.Points = append(m.Points, points)
	for x, c := range line {
		m.Points[y][x] = NewPoint(x, y, c)
	}
}

// At returns a point at specific coordinates.
func (m *Map2D) At(x, y int) *Point {
	if IsWithinMap(x, y, 0, 0, m.Width(), m.Height()) {
		return m.Points[y][x]
	}
	return nil
}

// Width returns the width of the map.
func (m *Map2D) Width() int {
	return len(m.Points[0])
}

// Height returns the height of the map.
func (m *Map2D) Height() int {
	return len(m.Points)
}

// ForAllPoints calls a function for each point in a map.
func (m *Map2D) ForAllPoints(f func(p *Point) bool) {
	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			if !f(m.Points[y][x]) {
				return
			}
		}
	}
}

// Next returns the next point in the given direction.
func (m *Map2D) Next(d Direction, p *Point) *Point {
	switch d {
	case DirectionUp:
		if p.Y > 0 {
			return m.Points[p.Y-1][p.X]
		}
	case DirectionDown:
		if p.Y < len(m.Points)-1 {
			return m.Points[p.Y+1][p.X]
		}
	case DirectionLeft:
		if p.X > 0 {
			return m.Points[p.Y][p.X-1]
		}
	case DirectionRight:
		if p.X < len(m.Points[p.Y])-1 {
			return m.Points[p.Y][p.X+1]
		}
	case DirectionUpRight:
		if p.Y > 0 && p.X < len(m.Points[p.Y-1])-1 {
			return m.Points[p.Y-1][p.X+1]
		}
	case DirectionDownRight:
		if p.Y < len(m.Points)-1 && p.X < len(m.Points[p.Y+1])-1 {
			return m.Points[p.Y+1][p.X+1]
		}
	case DirectionUpLeft:
		if p.X > 0 && p.Y > 0 {
			return m.Points[p.Y-1][p.X-1]
		}
	case DirectionDownLeft:
		if p.X > 0 && p.Y < len(m.Points)-1 {
			return m.Points[p.Y+1][p.X-1]
		}
	}
	return nil
}

func (m *Map2D) FindPath(start, end *Point, path rune) []*Point {
	prev := map[*Point]*Point{}
	dist := map[*Point]int{}
	pq := NewPriorityQueue[*Point]()
	pq.Push(start, 0)
	pathLen := 0
	for pq.Len() > 0 && pathLen == 0 {
		p, prio := pq.Pop()
		for _, d := range AllDirection {
			np := m.Next(d, p)
			if np == nil {
				continue
			}

			if np.C != path && np != end {
				continue
			}

			if _, ok := dist[np]; ok {
				continue
			}

			dist[np] = prio + 1
			prev[np] = p
			if np == end {
				pathLen = prio + 1
				break
			}
			pq.Push(np, prio+1)
		}
	}

	if pathLen == 0 {
		return nil
	}

	out := make([]*Point, pathLen)
	p := end
	for i := pathLen - 1; i >= 0; i-- {
		out[i] = p
		p = prev[p]
	}

	return out
}

// String implements the fmt.Stringer interface.
func (m *Map2D) String() string {
	var out strings.Builder
	for y := 0; y < len(m.Points); y++ {
		for x := 0; x < len(m.Points[y]); x++ {
			p := m.Points[y][x]
			out.WriteRune(p.C)
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func (m *Map2D) Debug(p *Point) {
	c := p.C
	p.C = '×'
	fmt.Println(m)
	p.C = c
	WaitInput()
}

// ManhattanDistance returns the ManhattanDistance between two points.
func ManhattanDistance(p1, p2 *Point) int {
	return Abs(p1.X-p2.X) + Abs(p2.Y-p1.Y)
}

// Vec represents a vector.
type Vec struct {
	U, V int
}

func NewVec(u, v int) Vec {
	return Vec{U: u, V: v}
}

// NewVecFromPoints returns a new vector from two points.
func NewVecFromPoints(a, b Point) Vec {
	return Vec{
		U: b.X - a.X,
		V: b.Y - a.Y,
	}
}

func (v Vec) Times(i int) Vec {
	v.U *= i
	v.V *= i
	return v
}

func IsWithinMap(x, y, minX, minY, maxX, maxY int) bool {
	if (x >= minX && x < maxX) && (y >= minY && y < maxY) {
		return true
	}

	return false
}

// Move represents a move for a point in a direction.
type Move struct {
	P *Point
	D Direction
}

func NewMove(p *Point, d Direction) Move {
	return Move{P: p, D: d}
}
