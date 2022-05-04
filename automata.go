package main

import (
	"math/rand"
)

// Automata ...
type Automata struct {
	N    int32
	grid [][][][]uint8
	ptr  uint8

	cube Object
}

// Init the automata
func (a *Automata) Init(N int) {
	a.N = int32(N)

	// Create grid
	a.grid = make([][][][]uint8, 2)
	for g := 0; g < 2; g++ {
		a.grid[g] = make([][][]uint8, a.N+2)
		for i := int32(0); i < a.N+2; i++ {
			a.grid[g][i] = make([][]uint8, a.N+2)
			for j := int32(0); j < a.N+2; j++ {
				a.grid[g][i][j] = make([]uint8, a.N+2)
			}
		}
	}

	// Fill grid
	s := a.N / 8
	m := a.N / 2
	var i, j, k int32
	for i = m - s; i <= m+s; i++ {
		for j = m - s; j <= m+s; j++ {
			for k = m - s; k <= m+s; k++ {
				if rand.Float32() < 0.7 {
					a.grid[a.ptr][i][j][k] = 1
				}
			}
		}
	}

	a.cube.Create(cubeVertices, cubeNormals)
}

func (a *Automata) countNeighbors(i, j, k int32) uint8 {
	var x, y, z int32

	n := -a.grid[a.ptr][i][j][k]
	for x = -1; x <= 1; x++ {
		for y = -1; y <= 1; y++ {
			for z = -1; z <= 1; z++ {
				n += a.grid[a.ptr][i+x][j+y][k+z]
			}
		}
	}

	return n
}

func (a *Automata) fracSimulate(i0, i1 int32, ch chan int) {
	var i, j, k int32

	for i = i0; i < i1; i++ {
		for j = 1; j <= a.N; j++ {
			for k = 1; k <= a.N; k++ {
				n := a.countNeighbors(i, j, k)

				a.grid[1-a.ptr][i][j][k] = 0
				if a.grid[a.ptr][i][j][k] == 0 {
					if n == 4 {
						// Birth
						a.grid[1-a.ptr][i][j][k] = 1
					}
				} else if n >= 4 && n <= 6 {
					// Survival
					a.grid[1-a.ptr][i][j][k] = 1
				}
			}
		}
	}

	ch <- 0
}

// Simulate the automata's next state
func (a *Automata) Simulate() {
	var i int32
	var step int32 = 20

	ch := make(chan int)

	for i = 1; i <= a.N; i += step {
		j := i + step
		if j > a.N {
			j = a.N
		}
		go a.fracSimulate(i, j, ch)
	}

	for i = 1; i <= a.N; i += step {
		<-ch
	}

	a.ptr = 1 - a.ptr
}

// Draw the current automata's state
func (a *Automata) Draw(prog *shaderProgram) {
	instances := []int32{}

	for i := int32(1); i <= a.N; i++ {
		for j := int32(1); j <= a.N; j++ {
			for k := int32(1); k <= a.N; k++ {
				if a.grid[a.ptr][i][j][k] > 0 {
					// Test surroundings
					d := false
					d = d || (a.grid[a.ptr][i-1][j][k] == 0)
					d = d || (a.grid[a.ptr][i+1][j][k] == 0)
					d = d || (a.grid[a.ptr][i][j-1][k] == 0)
					d = d || (a.grid[a.ptr][i][j+1][k] == 0)
					d = d || (a.grid[a.ptr][i][j][k-1] == 0)
					d = d || (a.grid[a.ptr][i][j][k+1] == 0)

					if d {
						instances = append(instances, []int32{i - a.N/2 + 1, j - a.N/2 + 1, k - a.N/2 + 1}...)
					}
				}
			}
		}
	}

	a.cube.UpdateInstances(instances)
	a.cube.Draw()
}
