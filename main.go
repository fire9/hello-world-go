package main

import (
	"fmt"

	"github.com/fire9/hello-world-go/doctor"
)

func main() {
	// var whatToSay string
	// whatToSay = doctor.Intro()
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
}
