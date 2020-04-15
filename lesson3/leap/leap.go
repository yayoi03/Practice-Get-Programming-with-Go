package main

import "fmt"

func main() {
	fmt.Println("2100年はうるう年？")
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
