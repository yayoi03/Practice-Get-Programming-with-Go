package main

import "fmt"

func main() {
	// var planets [8]string
	// planets[1] = "EARTH"
	// earth := planets[1]
	// fmt.Println(earth)

	// planets := [...]string{
	// 	"a",
	// 	"s",
	// 	"d",
	// }
	// fmt.Println(len(planets))

	//dwarfs := [5]string{"aa", "ss", "dd"}

	// for i := 0; i < len(dwarfs); i++ {
	// 	dwarf := dwarfs[i]
	// 	fmt.Println(i, dwarf)
	// }

	// for i, dwarf := range dwarfs {
	// 	fmt.Println(i, dwarf)
	// }

	// dwarfs2 := dwarfs
	// dwarfs[2] = "ddChanged"
	// fmt.Println(dwarfs)
	// fmt.Println(dwarfs2)

	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"

	for column := range board[1] {
		board[1][column] = "p"
	}

	fmt.Print(board)
}
