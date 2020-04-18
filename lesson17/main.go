package main

import (
	"fmt"
	"sort"
	"strings"
)

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	// planets := [...]string{
	// 	"aa",
	// 	"ss",
	// 	"dd",
	// 	"ff",
	// 	"gg",
	// 	"hh",
	// }
	// t := planets[0:1]
	// g := planets[1:3]
	// i := planets[3:6]
	// i2 := i[0:1]
	// i2[0] = "diuhsudhusd"
	// fmt.Println(t, g, i, i2)

	// planets := []string{"    aa", "ss   ", "dd   "}
	// hyperspace(planets)
	// fmt.Println(strings.Join(planets, ""))

	planets := []string{"aa", "dsd", "koo", "huh"}

	sort.StringSlice(planets).Sort()

	fmt.Println(planets)
}
