package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fire9/hello-world-go/doctor"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
	for {
		fmt.Print("->")
		userInput, _ := reader.ReadString('\n')

		// userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			// response := doctor.Response(userInput)
			// fmt.Println(response)
			fmt.Println(doctor.Response(userInput))
		}

	}

}
