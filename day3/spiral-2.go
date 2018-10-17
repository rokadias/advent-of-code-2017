package main

import (
	"fmt"
)

const input int = 10

type coor struct {
	x int
	y int
}

func main() {
	grid := make(map[coor]int)

	var x, y int = 0, 0
	grid[coor{x, y}] = 1
	x = 1
	max := 1
	i := max
	dir := 1
	currentValue := 0
	fmt.Println(grid[coor{0, 0}])

	for currentValue < input {
		currentValue := grid[coor{x - 1, y - 1}] + grid[coor{x, y - 1}] + grid[coor{x + 1, y - 1}] + grid[coor{x - 1, y}] + grid[coor{x, y}] + grid[coor{x + 1, y}] + grid[coor{x - 1, y + 1}] + grid[coor{x, y + 1}] + grid[coor{x + 1, y + 1}]
		grid[coor{x, y}] = currentValue

		if currentValue >= input {
			fmt.Println(currentValue)
			return
		}

		switch dir {
		case 0:
			x++
		case 1:
			y--
		case 2:
			x--
		case 3:
			y++
		}

		i--

		if i == 0 {
			if dir == 1 || dir == 3 {
				max++
			}
			i = max
			dir = (dir + 1) % 4
		}
	}

}
