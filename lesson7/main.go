package main

import (
	"fmt"
	"time"
)

func main() {
	// year := 2018
	// fmt.Printf("%T型: %v\n", year, year)

	// days := 365.2425
	// fmt.Printf("%T型: %[1]v \n", days)

	// var red uint8 = 255
	// red++
	// fmt.Println(red)

	// var number int8 = 127
	// number++
	// fmt.Println(number)

	future := time.Unix(12622780800, 0)
	fmt.Println(future)

}
