package routine

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func TestRoutine() {
	constructRoutine()
}

func constructRoutine() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()
	fmt.Println("Press any key or Q to quit")

	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'Q' || char == 'q' {
			break
		}

		keyPressChan <- char
	}

}

func listenForKeyPress() {

	for {
		key := <-keyPressChan
		fmt.Println("You pressed ", string(key))
	}
}
