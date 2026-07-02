package main

import (
	"encoding/json"
	"fmt"
)

type cource struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// Encodejson()
	Decodejson()
}

func Encodejson() {
	LearnableCources := []cource{
		{"GOLANG", 299, "youtube.com", "random11", []string{"Golang, backend, programming"}},
		{"Python", 199, "youtube.com", "random22", []string{"Python, backend, programming"}},
		{"Java", 399, "youtube.com", "random33", []string{"Java, backend, programming"}},
		{"JavaScript", 0, "youtube.com", "random123", nil},
	}

	jsondata, err := json.MarshalIndent(LearnableCources, "", "  ")
	if err != nil {
		panic(err)
	}

	println(string(jsondata))

}

func Decodejson() {
	jasondatafromweb := []byte(`[
	{
		"coursename": "GOLANG",
		"Price": 299,
		"website": "youtube.com",
		"tags": ["Golang, backend, programming"]
	},
	{
		"coursename": "Python",
		"Price": 199,
		"website": "youtube.com",
		"tags": ["Python, backend, programming"]
	}]`,
	)

	var googlecourses []cource

	validcheck := json.Valid(jasondatafromweb)

	if validcheck {
		fmt.Println("Json data was valid")

		json.Unmarshal(jasondatafromweb, &googlecourses)

		fmt.Printf("%#v\n", googlecourses)

	} else {
		fmt.Println("JSON data was not valid")
	}
}
