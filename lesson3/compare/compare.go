package main

import "fmt"

func main() {
	var age = 41
	var minor = age < 18
	fmt.Printf("%v才の僕は未成年か:%v？", age, minor)
}
