package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)

	//fmt.Println(id, "...snore...")
	c <- id
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	// go sleepyGopher()

	// time.Sleep(4 * time.Second)

	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	// for i := 0; i < 5; i++ {
	// 	gopherID := <-c
	// 	fmt.Println("gopher ", gopherID, "はスリープを終えました")
	// }

	// timeout := time.After(2 * time.Second)
	// for i := 0; i < 5; i++ {
	// 	select {
	// 	case gopherID := <-c:
	// 		fmt.Println("gopher ", gopherID, " はスリープを終えました。")
	// 	case <-timeout:
	// 		fmt.Println("忍耐の限度に達した！！！！！！！！！！")
	// 		return
	// 	}
	// }

	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
