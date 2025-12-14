package main

import (
    "fmt"
    "os"
)

func main() {
    envList := os.Environ()
    // выводим первые пять элементов
    for i := 0; i < len(envList); i++ {
        fmt.Println(envList[i])
    }
}