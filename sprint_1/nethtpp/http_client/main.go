package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		panic(err)
	}
	result := make([]byte, 512)
	n, err := response.Body.Read(result)
	response.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Print(n)
	fmt.Print(string(result))
}
