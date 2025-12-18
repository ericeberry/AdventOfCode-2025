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

	if c2.distance < c1.distance {
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

	var boxes []JBox
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

		boxes = append(boxes, JBox{x: x, y: y, z: z})
	}

	var connections Connections
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			d := calcDistance(boxes[i], boxes[j])
			connections = append(connections, Connection{a: boxes[i], b: boxes[j], distance: d})
		}
	}

	slices.SortFunc(connections, connectionComp)
	result := connectCircuits(connections, boxes)
	fmt.Println("result is ", result)
}

func connectCircuits(connections Connections, boxes []JBox) int {
	var circuits Circuits
	var connection Connection
	for i := 0; len(boxes) > 0; i++ {
		if i >= len(connections) {
			fmt.Println("error: not enough connections")
			return -1
		}
		connection = connections[i]
		a := -1
		b := -1
		for j, circuit := range circuits {
			if circuit.Contains(connection.a) {
				a = j
			}
			if circuit.Contains(connection.b) {
				b = j
			}
		}

		if a == -1 && b == -1 {
			circuit := NewCircuit()
			circuit.Add(connection.a)
			circuit.Add(connection.b)
			circuits = append(circuits, circuit)
			boxes = removeJBox(connection.a, boxes)
			boxes = removeJBox(connection.b, boxes)
		} else if a == -1 {
			circuits[b].Add(connection.a)
			boxes = removeJBox(connection.a, boxes)
		} else if b == -1 {
			circuits[a].Add(connection.b)
			boxes = removeJBox(connection.b, boxes)
		} else if a != b {
			circuits[a].Union(circuits[b])
			circuits = append(circuits[:b], circuits[b+1:]...)
		}
	}

	return connection.a.x * connection.b.x
}

func calcDistance(a JBox, b JBox) float64 {
	x2 := math.Pow(float64(a.x-b.x), 2)
	y2 := math.Pow(float64(a.y-b.y), 2)
	z2 := math.Pow(float64(a.z-b.z), 2)
	return math.Sqrt(x2 + y2 + z2)
}

func removeJBox(box JBox, boxes []JBox) []JBox {
	for i := range boxes {
		if box.x == boxes[i].x && box.y == boxes[i].y && box.z == boxes[i].z {
			return append(boxes[:i], boxes[i+1:]...)
		}
	}
	return boxes
}
