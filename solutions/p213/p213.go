package p213

import (
	"fmt"
	"math/rand"
	"strings"
)

type Simulation struct {
	Current    []int
	Next       []int
	SideLength int
	Generation int
}

func NewSimulation(sideLength int) Simulation {
	s := Simulation{
		SideLength: sideLength,
		Generation: 0,
	}

	s.Current = make([]int, sideLength*sideLength)
	s.Next = make([]int, sideLength*sideLength)

	for i := 0; i < len(s.Current); i++ {
		s.Current[i] = 1
	}

	return s
}

func (s *Simulation) Step() {
	for i := 0; i < len(s.Next); i++ {
		s.Next[i] = 0
	}

	for i, count := range s.Current {
		for flea := 0; flea < count; flea++ {
			s.Next[s.Move(i)]++
		}
	}

	s.Current, s.Next = s.Next, s.Current
}

func (s *Simulation) Move(index int) int {
	x := index % s.SideLength
	y := index / s.SideLength
	for { // rejection sampling
		xx, yy := x, y
		switch r := rand.Intn(4); r {
		case 0:
			xx++
		case 1:
			xx--
		case 2:
			yy++
		case 3:
			yy--
		}
		if xx >= 0 && xx < s.SideLength && yy >= 0 && yy < s.SideLength {
			return xx + yy*s.SideLength
		}
	}
}

func (s *Simulation) Reset() {
	s.Generation = 0
	for i := 0; i < len(s.Current); i++ {
		s.Current[i] = 1
	}
}

func (s *Simulation) Unoccupied() int {
	var res int
	for i := 0; i < len(s.Current); i++ {
		if s.Current[i] == 0 {
			res++
		}
	}
	return res
}

func (s *Simulation) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Generation: %d\n", s.Generation))
	b.WriteString(fmt.Sprintf("Unoccupied: %d\n", s.Unoccupied()))

	for i := 0; i < len(s.Current); i++ {
		b.WriteString(fmt.Sprintf("\t%d", s.Current[i]))
		if (i+1)%s.SideLength == 0 && i != len(s.Current)-1 {
			b.WriteString("\n")
		}
	}

	return b.String() // no trailing newline
}

func mean(x []int) float64 {
	var total float64 // no overflow
	for _, v := range x {
		total += float64(v)
	}
	return total / float64(len(x))
}

// Solve solves problem 213.
func Solve() float64 {
	var (
		sideLength = 30
		rings      = 50
		trials     = 10000
	)

	var results []int

	s := NewSimulation(sideLength)

	for t := 0; t < trials; t++ {
		for r := 0; r < rings; r++ {
			s.Step()
		}
		results = append(results, s.Unoccupied())
		s.Reset()
	}

	return mean(results)
}
