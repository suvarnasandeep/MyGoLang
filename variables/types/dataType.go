package types

import (
	"fmt"
	"sort"
)

type User struct {
	name    string
	version float32
	number  int
	isAuto  bool
}

type Animal struct {
	name     string
	sound    string
	noOfLegs int
}

// reciever function to type
func (a *Animal) says() {
	fmt.Printf("A %s says %s\n", a.name, a.sound)
}

func (a *Animal) numberOfLegs() {
	fmt.Printf("A %s has %d legs\n", a.name, a.noOfLegs)
}

var myArr [3]int

func TestDataType() {
	testStruct()
	testArray()
	testPointer()
	testSlices()
	testMap()

	fmt.Println(sumVarArg(1, 2, 3, 4))

	//reciever function test
	var dog Animal
	dog.name = "dog"
	dog.sound = " woof"
	dog.noOfLegs = 4
	dog.says()

	cat := Animal{
		name:     "cat",
		sound:    "meow",
		noOfLegs: 4,
	}
	cat.says()

	dog.numberOfLegs()
	cat.numberOfLegs()

}

func testStruct() {
	var user User
	user.name = "volvo"
	user.version = 2.1
	user.number = 11
	user.isAuto = true
	fmt.Println(user)

	user_1 := User{
		name:    "volkswagan",
		version: 2.0,
		number:  45,
		isAuto:  false,
	}
	fmt.Println(user_1)
}

func testArray() {
	myArr[0] = 1
	myArr[2] = 3

	fmt.Println(myArr)
	fmt.Printf("arr[0] = %d, arr[2] = %d \n", myArr[0], myArr[2])
}

func testPointer() {
	x := 5

	myPointer := &x
	fmt.Println("x value is : ", x)
	fmt.Println("myPointer value is : ", myPointer)

	*myPointer = 15
	fmt.Println("x value is : ", x)

	changePointerValue(&x)
	fmt.Println("x value is : ", x)
}

func testSlices() {
	var animal []string
	animal = append(animal, "cat")
	animal = append(animal, "dog")
	animal = append(animal, "fish")
	animal = append(animal, "cow")
	fmt.Print(animal)

	//iterarte
	for _, x := range animal {
		fmt.Println(x)
	}
	for i, x := range animal {
		fmt.Println(i, x)
	}

	fmt.Println(animal[0])
	fmt.Println(animal[0:2])
	fmt.Println(len(animal))
	fmt.Println("is it sorted ?", sort.StringsAreSorted(animal))
	sort.Strings(animal)
	fmt.Println("is it sorted ?", sort.StringsAreSorted(animal))

	fmt.Println(animal)
	animal = deleteFromSlice(animal, 1)
	fmt.Println(animal)
}

func testMap() {
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4

	fmt.Println(intMap)

	for key, val := range intMap {
		fmt.Println(key, val)
	}

	delete(intMap, "one")
	fmt.Println(intMap)

	el, ok := intMap["two"]
	if ok {
		fmt.Println(el, "present in map")
	} else {
		fmt.Println("not present")
	}

	intMap["two"] = 7
	fmt.Println(intMap)
}

func sumVarArg(num ...int) int {

	total := 0
	for _, x := range num {
		total = total + x
	}
	return total

}

func changePointerValue(num *int) {
	*num = 25
}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
