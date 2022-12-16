package main

import (
	"fmt"
	"gomod/test"
	"gomod/testone"
)

var one = "one"
var myVar = "package level var"

func main() {

	var blockVar = "block level var"
	fmt.Println("------------------------")
	fmt.Println(test.Morning)
	fmt.Println(testone.TestOneVar)
	testone.TestFunc()
	fmt.Println(one)
	otherFunc()
	fmt.Println("------------------------")

	test.PrintMe(myVar, blockVar)

}

func otherFunc() {
	//var one = "other one"
	fmt.Println(one)
}
