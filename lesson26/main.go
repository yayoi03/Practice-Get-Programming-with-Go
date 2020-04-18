package main

import (
	"fmt"
	"strings"
)

// type person struct {
// 	name, superpower string
// 	age              int
// }

// func birthday(p *person) {
// 	p.age++
// }

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew", int(1))
}

func main() {
	// answer := 42
	// fmt.Println(&answer)
	// adress := &answer
	// fmt.Println(*adress)

	// fmt.Printf("アドレスの型は%Tです", adress)

	// var ad *string
	// sc := "CJS"
	// ad = &sc
	// fmt.Println(*ad)
	// b := "CFB"
	// ad = &b
	// fmt.Println(*ad)

	// type person struct {
	// 	name, superpower string
	// 	age              int
	// }

	// timmy := &person{
	// 	name: "Timothy",
	// 	age:  10,
	// }

	// timmy.superpower = "flying"
	// fmt.Printf("%+v\n", timmy)

	// sp := &[3]string{"f", "i", "s"}
	// fmt.Println(sp[0])
	// fmt.Println(sp[1:2])

	// rebecca := person{
	// 	name:       "Rebecca",
	// 	superpower: "imagination",
	// 	age:        14,
	// }

	// birthday(&rebecca)

	// fmt.Printf("%+v\n", rebecca)

	// terry := &person{
	// 	name: "Terry",
	// 	age:  15,
	// }
	// terry.birthday()
	// fmt.Printf("%+v\n", terry)

	// nana := person{
	// 	name: "nana",
	// 	age:  11,
	// }

	// nana.birthday()
	// fmt.Printf("%+v\n", nana)

	// player := character{name: "sou"}
	// levelUp(&player.stats)
	// fmt.Printf("%+v\n", player.stats)

	// var board [8][8]rune
	// reset(&board)
	// fmt.Printf("%c", board[0][0])

	// planets := []string{
	// 	"a", "s", "d", "f", "g", "h", "j", "j", "k",
	// }

	// reclassify(&planets)
	// fmt.Println(planets)

	shout(martian{})
	shout(&martian{})

	pew := laser(2)
	shout(&pew)
}
