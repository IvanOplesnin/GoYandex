package main

import (
    "fmt"
    "time"
)

func main() {
    var today time.Time
    // допишите код
    // ...
    today = time.Now().Truncate(time.Hour * 24)

    fmt.Println(today)
}