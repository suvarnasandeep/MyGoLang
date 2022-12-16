package main

import (
	"bufio"
	"fmt"
	"myapp/routine"
	"os"
)

const prompt = "and press ENTER"

func main() {
	/*
		//one way - declare and then assign
		var firstNumber int
		firstNumber = 1

		//second way - declare type and var and assign value
		var secondNumber = 4

		//third way - declare name, assign value and go determines type
		thirdNumber := 6
	*/

	/*
		rand.Seed(time.Now().UnixNano())
		var firstNumber = rand.Intn(8) + 2
		var secondNumber = rand.Intn(8) + 2
		var thirdNumber = rand.Intn(8) + 2
		var answer = firstNumber*secondNumber - thirdNumber

		processResult(firstNumber, secondNumber, thirdNumber, answer)
	*/

	//types.TestDataType()
	//testString.TestSting()
	//consoleApp.TestConsoleApp()
	routine.TestRoutine()

}

func processResult(firstNumber int, secondNumber int, thirdNumber int, answer int) {
	//display welcome msg
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess the Number Game")
	fmt.Println("**********************")
	fmt.Println("Think a number between 1 to 10 ", prompt)
	reader.ReadString('\n')

	//take them through game
	fmt.Println("Multiply number by ", firstNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Multiply number by ", secondNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Devide the number you originally thought of ", prompt)
	reader.ReadString('\n')
	fmt.Println("Now subtract ", thirdNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("The answer is : ", answer)
}
