package main

import (
	"fmt"
)

func main() {
	printMessage("1")
	go printMessage("2")
	printMessage("3")
}

func printMessage(str string) {
	fmt.Printf(str)
}
