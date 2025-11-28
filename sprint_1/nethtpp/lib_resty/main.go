package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type User struct {
    ID       int     `json:"id"`
    Username string  `json:"username"`
    Email    string  `json:"email"`
}

func main() {
    var users []User
    url := "https://jsonplaceholder.typicode.com/users"

    // если выбрали resty, используйте SetResult(&users)
    // для получения результата сразу в виде массива
    // ...
	client := resty.New()
	resp, err := client.R().SetResult(&users).Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode() == 200 {
		for _, user := range users {
			fmt.Printf("%s ", user.Username)
		}
	}
}