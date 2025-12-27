package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
    currentTimeStr := "2021-09-19T15:59:41+03:00"
    currentTime, err := time.Parse(time.RFC3339, currentTimeStr)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(currentTime)

    now := time.Now()

    fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
    // добавьте проверки для методов After и Equal
    // ...

	fmt.Printf("Is %s after %s ? Answer: %t \n", now, currentTime, now.After(currentTime))
	fmt.Printf("Is %s equal %s ? Answer: %t \n", now, currentTime, now.Equal(currentTime))
}