package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// var curiosity struct {
	// 	lat  float64
	// 	long float64
	// }

	// curiosity.lat = -4.5895
	// curiosity.long = 137.4417

	// fmt.Println(curiosity.lat, curiosity.long)
	// fmt.Println(curiosity)

	// type location struct {
	// 	lat  float64
	// 	long float64
	// }

	// var spirit location
	// spirit.lat = 131.0
	// spirit.long = 323.0

	// var opportunity location
	// opportunity.lat = -1.9462
	// opportunity.long = 3.4442

	// fmt.Println(spirit, opportunity)

	// op2 := location{lat: 1.2, long: 3.4}

	// fmt.Println(op2)

	// op3 := location{1.2, 3.4}

	// fmt.Println(op3)
	// fmt.Printf("%+v\n", op3)

	// op4 := op3
	// op4.long += 1.0
	// fmt.Printf("%+v\n", op4)

	// type location struct {
	// 	name string
	// 	lat  float64
	// 	long float64
	// }

	// locations := []location{
	// 	{name: "gy", lat: 89.5, long: 32.5},
	// 	{name: "dgy", lat: 189.5, long: 232.5},
	// 	{name: "vgy", lat: 289.5, long: 332.5},
	// }

	// fmt.Println(locations)

	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity := location{1.1, 4.4}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))

}
