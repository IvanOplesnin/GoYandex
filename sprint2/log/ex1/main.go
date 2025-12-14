package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	newLog := log.New(&buf, "mylog: ", 0)
	newLog.Print("Hello, world!")
	newLog.Print("Goodbye")

	fmt.Print(&buf)
	// должна вывести
	// mylog: Hello, world!
	// mylog: Goodbye
}
