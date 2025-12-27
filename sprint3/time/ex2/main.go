package main

import (
    "fmt"
    "time"
)

func main() {
    currentTimeStr := "2021-09-19T15:59:41+03:00"
    t, _ := time.Parse("2006-01-02T15:04:05Z07:00", currentTimeStr) // "2006-01-02T15:04:05Z07:00"
	fmt.Print(t)
}