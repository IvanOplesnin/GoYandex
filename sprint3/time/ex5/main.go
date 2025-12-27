package main

import (
    "fmt"
    "time"
)

func main() {
    // допишите код здесь
    // birthday := time.Date(... , time.Local)
    // days := ...
    // fmt.Println(days)
	birthday := time.Date(2093, 11, 26, 0, 0, 0, 0, time.UTC)
	days := int(time.Until(birthday).Hours()/24)
	fmt.Println(days)
}