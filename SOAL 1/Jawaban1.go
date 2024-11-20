package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	x, y, r int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

func diDalam(lingkaran Lingkaran, titik Titik) bool {
	jarakKePusat := jarak(Titik{lingkaran.x, lingkaran.y}, titik)
	return jarakKePusat <= float64(lingkaran.r)
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titik Titik

	fmt.Scan(&lingkaran1.x, &lingkaran1.y, &lingkaran1.r)

	fmt.Scan(&lingkaran2.x, &lingkaran2.y, &lingkaran2.r)

	fmt.Scan(&titik.x, &titik.y)

	diDalam1 := diDalam(lingkaran1, titik)
	diDalam2 := diDalam(lingkaran2, titik)

	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
