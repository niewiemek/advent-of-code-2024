package helper

type Direction = int

const (
	INCREASING = iota
	DECREASING = iota
)

func AbsDiff(x int, y int) (diff int, direction Direction) {
	if x > y {
		diff = x - y
		direction = DECREASING
	} else {
		diff = y - x
		direction = INCREASING
	}

	return diff, direction
}
