package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// В високосном году 2016 было 366 дней.
	start := Date(2020, 1, 1)
	finish := Date(2020, 7, 31)
	day := math.Round(finish.Sub(start).Hours() / 24)
	fmt.Println(day) // 366
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
