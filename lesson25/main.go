package main

import "fmt"

type report struct {
	sol

	temperature

	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type sol int

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	return 5
}

func main() {
	//report := report{sol: 13, location: location{2.3, 1.5}, temperature: temperature{high: 1.0, low: -80.2}}
	//fmt.Printf("av:: %v", report.average())

	report := report{sol: 15}
	fmt.Println(report.sol.days(123))
}
