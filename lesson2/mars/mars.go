package main

import "fmt"

func main() {
	fmt.Print("火星で私の体重は")
	fmt.Print(149.0 * 0.3783)
	fmt.Println("火星で私の体重は", 149.0*0.3783)
	fmt.Printf("火星で私の体重は%vポンド", 149.0*0.3783)
	fmt.Printf("%-15v $%4v", "SpaceX", 94)
}
