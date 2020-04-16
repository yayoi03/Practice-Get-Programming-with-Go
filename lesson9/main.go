package main

import (
	"fmt"
)

func main() {
	// fmt.Println(`aa \n aa`)
	// fmt.Println("aa \n aa")
	// fmt.Println(`aa
	// dddd
	// ff`)

	// fmt.Printf("%vは%[1]T型です\n", "文字列リテラル")
	// fmt.Printf("%vは%[1]T型です\n", `文字列リテラル`)

	// var pi rune = 960
	// var alpha rune = 940
	// var omega rune = 969
	// var bang byte = 33

	// fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	// fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	// message := "shalom"
	// c := message[5]
	// fmt.Printf("%c \n", c)

	// message := "uv vagreangvbany fcnpr fgngvba"

	// for i := 0; i < len(message); i++ {
	// 	c := message[i]
	// 	if c >= 'a' && c <= 'z' {
	// 		c = c + 13
	// 		if c > 'z' {
	// 			c = c - 26
	// 		}
	// 	}
	// 	fmt.Printf("%c", c)
	// }

	// question := "cdskcks"
	// fmt.Println(len(question), "bytes")
	// fmt.Println(utf8.RuneCountInString(question), "runes")
	// c, size := utf8.DecodeRuneInString(question)
	// fmt.Printf("First rune: %c %v bytes", c, size)

	question := "cdskcks"

	for i, c := range question {
		fmt.Printf("%v %c \n", i, c)
	}
}
