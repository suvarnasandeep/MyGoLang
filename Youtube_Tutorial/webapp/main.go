package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const urll = "https://lco.dev"
const myUrl = "https://lco.dev:8000/learn?coursename=golang&paymentid=fdjgj5757mkm"

func main() {
	fmt.Println("This is my web app...!!!")

	//getWebRequest()

	handleWebURL()
}
func handleWebURL() {
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)
	fmt.Println("Scheme : ", result.Scheme)
	fmt.Println("Host :", result.Host)
	fmt.Println("Path :", result.Path)
	fmt.Println("Raw Query", result.RawQuery)
	fmt.Println("Port :", result.Port())

	qparams := result.Query()
	fmt.Println("query params : ", qparams)
	fmt.Println(qparams["coursename"])

	//construct URL
	partOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/code",
		RawPath: "user=sandeep",
	}

	anotherURL := partOfURL.String()
	fmt.Println(anotherURL)

}

func getWebRequest() {
	//web request
	response, err := http.Get(urll)
	errNilCheck(err)

	fmt.Printf("Response type is %T\n", response)
	//fmt.Println(response)
	defer response.Body.Close()

	dataByte, err := ioutil.ReadAll(response.Body)
	errNilCheck(err)

	content := string(dataByte)
	fmt.Println(content)
}

func errNilCheck(err error) {
	if err != nil {
		panic(err)
	}
}
