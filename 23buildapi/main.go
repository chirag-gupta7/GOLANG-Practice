package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for Cource -File
type Cource struct {
	CourceName  string  `json:"coursename"`
	CourceId    string  `json:"courseid"`
	CourcePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake Database
var cources []Cource

func (c *Cource) IsEmpty() bool {
	return c.CourceName == "" && c.CourceId == ""
}

func main() {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API</h1>"))

}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Cources")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources)
}
