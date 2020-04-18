package main

import (
	"fmt"
)

type kelvin float64
type sensor func() kelvin

// func fakeSensor() kelvin {
// 	return kelvin(rand.Intn(151) + 150)
// }

func realSensor() kelvin {
	return 0
}

// func measureTemperature(samples int, sensor func() kelvin) {
// 	for i := 0; i < samples; i++ {
// 		k := sensor()
// 		fmt.Printf("%vo k\n", k)
// 		time.Sleep(time.Second)
// 	}
// }

// var f = func() {
// 	fmt.Println("Dress up for the masquerade.")
// }

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	// sensor := fakeSensor
	// fmt.Println(sensor())

	// sensor = realSensor
	// fmt.Println(sensor())

	// measureTemperature(3, fakeSensor)

	// f()

	// func() {
	// 	fmt.Println("2 Dress up for the masquerade.")
	// }()

	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}
