package main

import (
	"log"

	"github.com/TMaulana26/GolangChallengeEP1/app"
)

func main() {
	log.Println(app.Multiply(7.8, 2))

	if isValid, err := app.IsInteger("1"); nil != err {
		log.Println(isValid, err.Error())
	} else {
		log.Println(isValid)
	}

	log.Println(app.IntToDayName(4))

	log.Println(app.DayNameToInt("minggu"))

	var kosong string
	app.IsiHello(&kosong)
	log.Println(kosong)
}
