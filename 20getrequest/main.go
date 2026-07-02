package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest("http://localhost:8000/get")
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(url string) string {

	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	var responceString strings.Builder

	content, err := io.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}
	bytecount, err := responceString.Write(content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Byte count:", bytecount)
	fmt.Println("Respnce string: ", responceString.String())

	return responceString.String()

}

func PerformPostJsonRequest() {

	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)
	responce, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	content, err := io.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("Respnce from server: ", string(content))
}

func PerformPostFormRequest() {

	const myurl = "http://localhost:8000/postform"
	data := url.Values{}

	data.Add("firstname", "John")
	data.Add("lastname", "Doe")
	data.Add("email", "random@go.dev")

	responce, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()
	content, err := io.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("Respnce from server: ", string(content))
}
