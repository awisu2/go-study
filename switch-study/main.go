package main

import "log"

func main () {
	y := sample(1)
	log.Println(y)

	y = sample(2)
	log.Println(y)

	y = sample(4)
	log.Println(y)

	y = sample(100)
	log.Println(y)
}

func sample(x int) int {
	y := 0
	switch x {
	case 2, 4:
		y = 2
	case 1:
		y = 1
	default:
		y = 0
	}
	return y
}
