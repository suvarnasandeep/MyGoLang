package testString

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	name      string
	age       int
	favNumber float64
}

// sample structure
type MyStruct struct {
	name    string
	dob     int
	isHuman bool
	number  float64
}

func TestSting() {
	mainTest()
}

func mainTest() {

	//printing struct
	myStruct := MyStruct{
		name:    "sandeep",
		dob:     1990,
		isHuman: true,
	}
	fmt.Println(myStruct)

	reader = bufio.NewReader(os.Stdin)

	//userName := readString("What is your name?")
	//age := readInt("What is your age?")

	var user User
	user.name = readString("What is your name?")
	user.age = readInt("What is your age?")
	user.favNumber = readFloat("What is your fav number")

	//fmt.Println("Your name is"+readString("What is your name")+". Your age is", readInt("What is your age"))
	//fmt.Println(fmt.Sprintf("Your name is %s. Your age is %d", userName, age))
	//fmt.Printf("Your name is %s. Your age is %d\n", userName, age)
	fmt.Printf("Your name is %s. Your age is %d. Your fav number is %.2f\n",
		user.name,
		user.age,
		user.favNumber)
}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userAge, _ := reader.ReadString('\n')
		userAge = strings.Replace(userAge, "\r\n", "", -1)
		userAge = strings.Replace(userAge, "\n", "", -1)

		age, err := strconv.Atoi(userAge)

		if err != nil {
			fmt.Println("Enter proper number..!!")
		} else {
			return age
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userAge, _ := reader.ReadString('\n')
		userAge = strings.Replace(userAge, "\r\n", "", -1)
		userAge = strings.Replace(userAge, "\n", "", -1)

		age, err := strconv.ParseFloat(userAge, 64)

		if err != nil {
			fmt.Println("Enter proper age..!!")
		} else {
			return age
		}
	}
}
