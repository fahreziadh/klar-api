package utils

import (
	"fmt"
	"math"
	"time"
)

func RoundTime(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}

// func getTotalDayInMonth(t string) {
// 	layOut := "02/01/2006 15:04:05" // dd/mm/yyyy hh:mm:ss
// 	future, err := time.Parse(layOut, "07/05/2018 15:12:10")
// 	if err != nil {
// 		panic(err)
// 	}
// 	diff := time.Until(future)
// 	fmt.Println("Now : ", time.Now())

// }

func GetDetail() {
	layOut := "02/01/2006 15:04:05" // dd/mm/yyyy hh:mm:ss
	future, err := time.Parse(layOut, "07/05/2022 15:12:10")

	if err != nil {
		panic(err)
	}
	diff := time.Until(future)

	fmt.Println("Now : ", time.Now())
	fmt.Println("Time reach the future date : ", future)

	fmt.Println("Raw : ", diff)
	fmt.Println("Hours : ", diff.Hours())
	fmt.Println("Minutes : ", diff.Minutes())
	fmt.Println("Seconds : ", diff.Seconds())
	fmt.Println("Nano seconds : ", diff.Nanoseconds)

	// get day, month and year

	fmt.Println("Days : ", RoundTime(diff.Seconds()/86400))
	fmt.Println("Weeks : ", RoundTime(diff.Seconds()/604800))
	fmt.Println("Months : ", RoundTime(diff.Seconds()/2600640))
	fmt.Println("Years : ", RoundTime(diff.Seconds()/31207680))
}
