package consoleApp

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func TestConsoleApp() {
	mainTest()
}

func mainTest() {
	/*
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("->")
			userInput, _ := reader.ReadString('\n')

			userInput = strings.Replace(userInput, "\n", "", -1)

			if userInput == "quit" {
				break
			} else {
				fmt.Println(userInput)
			}
		}
	*/

	/*
		err := keyboard.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			_ = keyboard.Close()
		}()

		fmt.Println("Press any key on keyboars, Press ESC to quit")

		for {
			char, key, err := keyboard.GetSingleKey()

			if err != nil {
				log.Fatal(err)
			}

			if key != 0 {
				fmt.Println("You pressed ", char, key)
			} else {
				fmt.Println("You pressed ", char)
			}

			if key == keyboard.KeyEsc {
				break
			}

		}
		fmt.Println("Application is exiting...!!!")
	*/

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	cofees := make(map[int]string)
	cofees[1] = "Cappucino"
	cofees[2] = "Latte"
	cofees[3] = "Expresso"
	cofees[4] = "Mocha"

	fmt.Println("Menu")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Expresso")
	fmt.Println("4 - Mocha")
	fmt.Println("Q - Quit")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}
		if char == 'Q' || char == 'q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", cofees[i]))

	}
	fmt.Println("Application is exiting...!!!")
}
