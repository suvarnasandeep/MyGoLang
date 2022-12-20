package scope

import (
	"fmt"
	"log"

	"myapp/scope/testscope"
	"myapp/scope/testscopeone"
)

var one = "one"
var myVar = "package level var"

var employees = []testscope.Employee{
	{FirstName: "sandeep", LastName: "Suvarna", Salary: 50000, FullTime: true},
	{FirstName: "Niriksha", LastName: "Anil", Salary: 60000, FullTime: true},
	{FirstName: "Anup", LastName: "Kumar", Salary: 70000, FullTime: true},
	{FirstName: "Rakshi", LastName: "Shetty", Salary: 80000, FullTime: false},
}

func TestScope() {
	var blockVar = "block level var"
	fmt.Println("------------------------")
	fmt.Println(testscope.Morning)
	fmt.Println(testscopeone.TestOneVar)
	testscopeone.TestFunc()
	fmt.Println(one)
	otherFunc()
	fmt.Println("------------------------")

	testscope.PrintMe(myVar, blockVar)
	fmt.Println("------------------------")

	//another example of export
	myStaff := testscope.Office{
		AllStaff: employees,
	}

	log.Println(myStaff.Show())
	log.Println(myStaff.OverPaid())
}

func otherFunc() {
	//var one = "other one"
	fmt.Println(one)
}
