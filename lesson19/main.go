package main

import (
	"fmt"
)

func main() {
	// temperature := map[string]int{
	// 	"Earth": 15,
	// 	"Mars":  -65,
	// }
	// temp := temperature["Earth"]
	// fmt.Printf("Earth: %v", temp)

	// temperature["Earth"] = 16
	// temperature["Venus"] = 464
	// fmt.Println(temperature)

	// t2 := temperature
	// t2["Earth"] = 99
	// fmt.Println(temperature)

	temperatures := []float64{
		28.0, 32.0, 31.0, 32.0, 29.0,
	}

	// frequency := make(map[float64]int)

	// for _, t := range temperatures {
	// 	frequency[t]++
	// }

	// for t, num := range frequency {
	// 	fmt.Printf("%+.2fの出現は%d回です\n", t, num)
	// }

	// groups := make(map[float64][]float64)

	// for _, t := range temperatures {
	// 	g := math.Trunc(t/10) * 10
	// 	groups[g] = append(groups[g], t)
	// }

	// for g, temperatures := range groups {
	// 	fmt.Printf("%v: %v \n", g, temperatures)
	// }

	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	if set[32.0] {
		fmt.Println("member of SET")
	}

	fmt.Println(set)
}
