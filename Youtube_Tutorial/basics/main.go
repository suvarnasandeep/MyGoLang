package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

/*
https://www.youtube.com/watch?v=62qGe9yhiJI&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=3

env var :
	go env
build:
	go build : builds for default OS
	GOOS="windows" go build
*/

type User struct {
	name     string
	lastname string
	age      int
}

func main() {
	//time
	printSeparator()
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2021, time.May, 30, 04, 30, 23, 222, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

	//user input
	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter pizza rating between 1 to 5..!!")
		userInput, _ := reader.ReadString('\n')
		fmt.Println("Thanks for rating, ", userInput)
		fmt.Printf("Rating type : %T\n", userInput)

		ratingInt, _ := strconv.Atoi(strings.TrimSpace(userInput))
		fmt.Println("rating is ", ratingInt)
		fmt.Printf("Rating type : %T\n", ratingInt)

		ratingFloat, _ := strconv.ParseFloat(strings.TrimSpace(userInput), 64)
		fmt.Println("rating is ", ratingFloat)
		fmt.Printf("Rating type : %T\n", ratingFloat)
	*/

	/*
		//pointers
		number := 25
		var ptr *int
		fmt.Println(ptr)

		ptr = &number
		fmt.Println(ptr)
		fmt.Println(*ptr)
		fmt.Println(&ptr)
	*/

	/*
		//slices
		//1st way of declaring
		var fruitList = []string{"mango", "banana", "peach", "apple", "chiku"}
		fmt.Printf("Type of fruitList : %T\n", fruitList)
		fmt.Println(fruitList)
		fruitList = append(fruitList, "grapes")
		fmt.Println(fruitList)

		fruitList = append(fruitList[1:])
		fmt.Println(fruitList)

		fruitList = append(fruitList[:3])
		fmt.Println(fruitList)

		fruitList = append(fruitList[1:2])
		fmt.Println(fruitList)

		//2nd way of declaring
		var scores = make([]int, 4)
		scores[0] = 111
		scores[1] = 222
		scores[2] = 333
		scores[3] = 444
		//scores[4] = 444

		scores = append(scores, 555, 666)
		fmt.Println(scores)
		fmt.Println(sort.IntsAreSorted(scores))

		//remove value from slice
		var fruitList1 = []string{"mango", "banana", "peach", "apple", "chiku"}
		fmt.Println(fruitList1)
		index := 2
		fruitList1 = append(fruitList1[:index], fruitList1[index+1:]...)
		fmt.Println(fruitList1)
	*/

	/*
		//map example
		languanges := make(map[string]string)
		languanges["JS"] = "javascript"
		languanges["RB"] = "ruby"
		languanges["PY"] = "python"
		fmt.Println(languanges)
		fmt.Println(languanges["JS"])
		delete(languanges, "JS")
		fmt.Println(languanges)

		for key, val := range languanges {
			fmt.Println(key, " : ", val)
		}
	*/

	/*
		//struct
		user := User{
			name:     "sandeep",
			lastname: "suvarna",
			age:      30,
		}
		fmt.Println(user)

		user1 := User{"niri", "suvarna", 26}
		fmt.Println(user1)
		fmt.Printf("user1 details are %+v\n", user1)
	*/

	/*
		//switch case
		rand.Seed(time.Now().UnixNano())
		diceNumber := rand.Intn(6) + 1
		fmt.Println("Dice number is ", diceNumber)
		switch diceNumber {
		case 1:
			fmt.Println("Dice value is 1")
		case 2:
			fmt.Println("Dice value is 2")
		case 3:
			fmt.Println("Dice value is 3")
		case 4:
			fmt.Println("Dice value is 4")
			fallthrough
		case 5:
			fmt.Println("Dice value is 5")
		case 6:
			fmt.Println("Dice value is 6")
		default:
			fmt.Println("default case")
		}
	*/
	/*
		//defer
		printSeparator()
		defer fmt.Println("This is defer")
		defer fmt.Println("This is defer1")

		defer func() {
			fmt.Println("this is defer function")
		}()

		myDeferTest()

		fmt.Println("heloooooo")

		defer myDeferTest()
		fmt.Println("this is last line")
	*/

	/*
		//file handling
		printSeparator()
		content := "This is my go file content"
		file, err := os.Create("./myGoFile.txt")

		// if err != nil {
		// 	panic(err)
		// }
		checkNilError(err)

		length, err := io.WriteString(file, content)
		checkNilError(err)
		fmt.Println("length is ", length)
		defer file.Close()

		readFile("./myGoFile.txt")
		printSeparator()
	*/

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilError(err)

	fmt.Println(string(dataByte))
}

func myDeferTest() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func printSeparator() {
	fmt.Println("------------------------------")
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
