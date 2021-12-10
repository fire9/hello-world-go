package main

import "fmt"

func main() {
	// var whatToSay string
	// whatToSay = "Hellow World again!"
	whatToSay := "Hellow World again!"
	fmt.Println("Hello, World!")
	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
