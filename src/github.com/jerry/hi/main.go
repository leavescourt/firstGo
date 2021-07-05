package main

import (
	"fmt"
	"github.com/jerry/dates"
	"github.com/jerry/greeting"
)

func main() {
	greeting.Hello()
	days := 12
	weeks := dates.DaysToWeeks(days)
	fmt.Printf("%d days is %.2f weeks\n", days, weeks)
}
