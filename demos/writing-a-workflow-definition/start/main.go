package main

import (
	"app"
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	greeting := app.GreetSomeone(name)
	fmt.Println(greeting)
}
