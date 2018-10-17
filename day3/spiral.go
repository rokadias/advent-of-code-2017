package main

import (
	"fmt"
	"math"
)

const input int = 289326

func showResult(x, y int) {
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func main() {
	var i int = 1

	for ; (i * i) < input; i = i + 2 {
	}

	squared := i * i

	x, y := (i / 2), (i / 2)
	fmt.Println(x, y)
	if squared-i+1 > input {
		x = -1 * x
	} else {
		x -= squared - input
		showResult(x, y)
		return
	}

	if squared-(2*(i-1)) > input {
		y = -1 * y
	} else {
		y -= squared - (i - 1) - input
		showResult(x, y)
		return
	}

	if squared-(3*(i-1)) > input {
		x = -1 * x
	} else {
		x += squared - (2 * (i - 1)) - input
		showResult(x, y)
		return
	}

	y += squared - (3 * (i - 1)) - input
	showResult(x, y)
	return
}
