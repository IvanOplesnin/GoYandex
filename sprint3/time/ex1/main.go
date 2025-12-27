package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Print(now.Format(time.RFC1123))
}