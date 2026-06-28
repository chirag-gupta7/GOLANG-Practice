package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://chirag-portfolio-new.vercel.app"

func main() {

	responce, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Responce Type : %T", responce)

	defer responce.Body.Close()

	urldata, err := io.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}

	content := string(urldata)

	fmt.Println(content) /*  */

}
