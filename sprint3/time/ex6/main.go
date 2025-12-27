package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(2 * time.Second)
	t := time.Now()
	for true {
		_ = <- ticker.C
		fmt.Printf("Duration: %d\n", int(time.Since(t).Seconds()))
	}
}