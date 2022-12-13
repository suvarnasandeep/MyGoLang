package main

import (
	"fmt"
	"gomod/test"
)

var one = "one"

func main() {

	fmt.Println(test.Morning)
	fmt.Println(one)
	otherFunc()

}

func otherFunc() {
	//var one = "other one"
	fmt.Println(one)
}
