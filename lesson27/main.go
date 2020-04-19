package main

import (
	"fmt"
	"sort"
)

func sortString(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

var v interface{}

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "未設定"
	}
	return fmt.Sprintf("%d", n.value)
}

func main() {
	// var fn func(a, b int) int

	// fmt.Println(fn == nil)

	// food := []string{"oni", "car", "cel"}
	// sortString(food, nil)
	// fmt.Println(food)

	// var p *int
	// v = p
	// fmt.Printf("%T %v %v \n", v, v, v == nil)

	// fmt.Printf("%#v \n", v)

	n := newNumber(42)
	fmt.Println(n)

	e := number{}
	fmt.Println(e)
}
