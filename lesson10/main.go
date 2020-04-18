package main

import (
	"fmt"
)

func main() {
	// age := 41
	// marsAge := float64(age)
	// marsDays := 687.0
	// earthDays := 365.2425
	// marsAge = marsAge * earthDays / marsDays
	// fmt.Println("私は火星では", marsAge, "歳です")

	// var pi rune = 960
	// var bang byte = 33
	// fmt.Print(string(pi), string(bang))

	// countdown := 10
	// str := "発射まで" + strconv.Itoa(countdown) + "秒。"
	// fmt.Println(str)

	// launch := false
	// launchText := fmt.Sprintf("%v", launch)
	// fmt.Println("Ready for launch:", launchText)

	// var yesNo string
	// if launch {
	// 	yesNo = "yes"
	// } else {
	// 	yesNo = "no"
	// }

	// fmt.Println("Ready for launch:", yesNo)

	yesNo := "no"
	launch := (yesNo == "yes")
	fmt.Println("Ready for launch:", launch)

}
