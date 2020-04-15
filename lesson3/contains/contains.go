package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("君は暗い洞窟の中にいる")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("洞窟を出る：", exit)
}
