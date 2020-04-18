package main

import (
	"fmt"
	"strings"
	"time"
)

var t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(3))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

type stardater interface {
	YearDay() int
	Hour() int
}

func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	t = martian{}
	fmt.Println(t.talk())

	//laser(1)でint型の1をlaser型の1にしている
	t = laser(1)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(2))

	// s := starship{laser(3)}
	// fmt.Println(s.talk())
	// shout(s)

	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f ", stardate(day))

	s := sol(1422)
	fmt.Printf("%.1f\n", stardate(s))

	cu := location{23.4, 4.53}
	fmt.Println(cu)
}
