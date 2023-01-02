package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func (u *User) print() {
	u.age = 40
	fmt.Println(u.name, u.lastname, u.age)
}

// interface example
type Animal interface {
	says() string
	hasLegs() int
}

type Dog struct {
	name string
	legs int
}

type Cat struct {
	name    string
	hasTail bool
}

func (d *Dog) says() string {
	return "woof"
}

func (d *Dog) hasLegs() int {
	return 4
}

// json support
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog`
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

		color := "green"
		fmt.Println("Color is ", color)
		callPointerFunc(&color)
		fmt.Println("Color is ", color)
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
		fmt.Println("---------------")
		user.print()
		fmt.Println(user.age)
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

	/*
		//interface
		dog := Dog{
			name: "rocky",
			legs: 4,
		}

		printDog(&dog)
	*/

	//readAndWriteJson()

	chanelExample()
}

func chanelExample() {
	fmt.Println("First line")
	c := make(chan int)
	defer close(c)

	go printGoRoiutine(c)

	fmt.Println("Second line")
	c <- 10
	//go printGoRoiutine(c)
	//c <- 20
	fmt.Println("Last line")

	myChan := make(chan string)

	go myFunc(myChan)

	for {
		res, ok := <-myChan
		if ok == false {
			fmt.Println("channel close ", ok)
			break
		}
		fmt.Println("channel open ", res, ok)
	}

}

func myFunc(myChan chan string) {

	for i := 0; i < 4; i++ {
		myChan <- "sandeep"
	}
	close(myChan)
}

func printGoRoiutine(c chan int) {
	fmt.Println("Data in the chanel is :", <-c)
}

func readAndWriteJson() {
	myJson := `
	[
		{
			"first_name": "Tony",
			"last_name": "StarK",
			"hair_color":"grey",
			"has_dog":true
		},
		{
			"first_name": "Bruce",
			"last_name": "Banner",
			"hair_color":"green",
			"has_dog":false
		}
	]
	`
	//write json to struct
	var unMarshalled []Person

	error := json.Unmarshal([]byte(myJson), &unMarshalled)
	if error != nil {
		log.Println(error)
	}
	log.Printf("Unmarshalled : %v\n", unMarshalled)

	//write struct to json

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Clark"
	m1.LastName = "Kent"
	m1.HairColor = "black"
	m1.HasDog = true

	var m2 Person
	m2.FirstName = "Clark"
	m2.LastName = "Kent"
	m2.HairColor = "black"
	m2.HasDog = true

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}

func printDog(a Animal) {
	fmt.Println("dog says : ", a.says(), "and has ", a.hasLegs(), "legs")
}

func callPointerFunc(s *string) {
	fmt.Println(s)
	newColor := "Red"
	*s = newColor
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
