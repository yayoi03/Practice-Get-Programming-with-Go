package main

import "fmt"

// func dump(label string, slice []string) {
// 	fmt.Printf("%v: length %v, capacity %v %v \n", label, len(slice), cap(slice), slice)
// }

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	// dwarfs := []string{"a", "s", "sds", "hugh", "ft"}
	// // dwarfs = append(dwarfs, "ffff", "oko")
	// // fmt.Println(dwarfs)

	// dump("dwarfs", dwarfs)

	// dwarfs2 := append(dwarfs, "xauh")
	// dump("d2", dwarfs2)

	// dwarfs3 := append(dwarfs2, "d3", "gg", "hu")
	// dump("d3", dwarfs3)

	// dwarfs3[1] = "testaaa"

	// dump("dwarfs", dwarfs)

	// dump("d2", dwarfs2)

	// dump("d3", dwarfs3)

	// planets := []string{
	// 	"aa", "sa", "hu", "ijj",
	// 	"ojj", "baa", "nsa", "nhu",
	// }

	// t := planets[0:4:4]
	// w := append(t, "流石")

	// fmt.Println(w)

	// fmt.Println(planets)

	// dwarfs := make([]string, 0, 10)
	// dwarfs = append(dwarfs, "jij", "ftft", "jjji")
	// fmt.Println(dwarfs)

	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	planets := []string{"hihi", "huhu", "oov", "jojojjso"}
	newPlanets := terraform("Super", planets...)
	fmt.Println(newPlanets)
}
