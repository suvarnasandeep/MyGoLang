package main

import (
	"bufio"
	"fmt"
	"myGoApp/doctor"
	"os"
	"strings"
)

func main() {
	firstMethod()

	//read content from user
	reader := bufio.NewReader(os.Stdin)

	intro := doctor.Intro()
	fmt.Println(intro)

	for {
		fmt.Print("You : ")
		//process when enter pressed
		userInput, _ := reader.ReadString('\n')

		//for windows
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		//for mac and linux
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		}
		fmt.Println("Eliza : " + doctor.Response(userInput))
	}

}

func firstMethod() {
	var tempVar string
	tempVar = "Sandeep"

	tempVar2 := "Suvarna."

	fmt.Println("*******************************")
	secondMethod("Hi!! This is  " + tempVar + " " + tempVar2)
	fmt.Println("*******************************")
}

func secondMethod(whatToPrint string) {
	fmt.Println(whatToPrint)
}
