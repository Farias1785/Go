package main

import (
	"fmt"
	"time"
)

func main() {
	atual()
}

func atual() {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	startday := firstOfMonth
	fmt.Println(firstOfMonth)
	fmt.Println(lastOfMonth)
	fmt.Println(startday)
}
