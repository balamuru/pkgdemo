package main

import (
	"fmt"
	"github.com/balamuru/pkgdemo/greetings"
)

func main() {
	message := greetings.Hello("VB")
	fmt.Println(message)
}
