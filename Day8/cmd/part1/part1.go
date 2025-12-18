package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const maxConnections = 1000

type JBox struct {
	x int
	y int
	z int
}

type Connection struct {
	a        JBox
	b        JBox
	distance float64
}

type Connections []Connection

func connectionComp(c1 Connection, c2 Connection) int {
	if c1.distance < c2.distance {
		return -1
	}

	if c1.distance < c2.distance {
		return 1
	}

	return 0
}

type Circuit map[JBox]struct{}

func NewCircuit() Circuit      { return make(Circuit) }
func (c Circuit) Add(box JBox) { c[box] = struct{}{} }
func (c Circuit) Contains(box JBox) bool {
	_, exists := c[box]
	return exists
}
func (c Circuit) Union(other Circuit) {
	for element := range other {
		c[element] = struct{}{}
	}
}

type Circuits []Circuit

func circuitsComp(c1 Circuit, c2 Circuit) int {
	if len(c1) > len(c2) {
		return -1
	}

	if len(c1) < len(c2) {
		return 1
	}

	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("error opening file input.txt:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var points []JBox
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("error parsing float:", err)
			return
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("error parsing float:", err)
			return
		}

		z, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("error parsing float:", err)
			return
		}

		points = append(points, JBox{x: x, y: y, z: z})
	}

	var connections Connections
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d := calcDistance(points[i], points[j])
			connections = append(connections, Connection{a: points[i], b: points[j], distance: d})
		}
	}

	slices.SortFunc(connections, connectionComp)
	result := connectCircuits(connections)
	fmt.Println("result is ", result)
}

func connectCircuits(connections Connections) int {
	var circuits Circuits
	for i := 0; i < maxConnections; i++ {
		if i >= len(connections) {
			fmt.Println("error: not enough connections")
			return -1
		}
		a := -1
		b := -1
		for j, circuit := range circuits {
			if circuit.Contains(connections[i].a) {
				a = j
			}
			if circuit.Contains(connections[i].b) {
				b = j
			}
		}

		if a == -1 && b == -1 {
			circuit := NewCircuit()
			circuit.Add(connections[i].a)
			circuit.Add(connections[i].b)
			circuits = append(circuits, circuit)
		} else if a == -1 {
			circuits[b].Add(connections[i].a)
		} else if b == -1 {
			circuits[a].Add(connections[i].b)
		} else if a != b {
			circuits[a].Union(circuits[b])
			circuits = append(circuits[:b], circuits[b+1:]...)
		}
	}

	slices.SortFunc(circuits, circuitsComp)
	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func calcDistance(a JBox, b JBox) float64 {
	x2 := math.Pow(float64(a.x-b.x), 2)
	y2 := math.Pow(float64(a.y-b.y), 2)
	z2 := math.Pow(float64(a.z-b.z), 2)
	return math.Sqrt(x2 + y2 + z2)
}
