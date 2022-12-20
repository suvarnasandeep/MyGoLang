package testscope

import "fmt"

var Morning = "sandeep"
var PkgVar = "Package var"

func PrintMe(myVar, blockVar string) {
	fmt.Println("I am printing in test...!!!")

	fmt.Println("myVar : ", myVar+" blockVar : ", blockVar)
	fmt.Println(`pkgVar : `, PkgVar+
		` print function : `)
}
