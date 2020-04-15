package main

import (
	"math/rand"
)

func main() {
	switch num := rand.Intn(10); num {
	case 0:
		println("SA")
	case 1:
		println("SX")
	case 2:
		println("VG")
	default:
		println("Random spaceline #", num)
	}
}
