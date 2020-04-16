package main

import "fmt"

func main() {
	// const lightSpeed = 299792
	// const secondsPerDay = 86400

	// var distance int64 = 41.3e12
	// fmt.Println("アルファ・ケンタウリまでの距離は、", distance, "km")

	// days := distance / lightSpeed / secondsPerDay
	// fmt.Println("光速で", days, "日かかる")

	// lightSpeed := big.NewInt(299792)
	// secondsPerDay := big.NewInt(86400)
	// distance := new(big.Int)
	// distance.SetString("24000000000000", 10)
	// fmt.Println("アンドロメダ銀河までの距離は、", distance, "km")

	// seconds := new(big.Int)
	// seconds.Div(distance, lightSpeed)

	// days := new(big.Int)
	// days.Div(seconds, secondsPerDay)
	// fmt.Println("光の速度で", days, "日かかる")

	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const days = distance / lightSpeed / secondsPerDay
	fmt.Println(days)
}
