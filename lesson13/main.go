package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	// type celsius float64

	// var temperature celsius = 20
	// fmt.Println(temperature)

	// var warmUp float64 = 10
	// fmt.Println(temperature + celsius(warmUp))

	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Print(k, "kは", c, "cです\n")
	c2 := k.celsius()
	fmt.Print(c2)

}
